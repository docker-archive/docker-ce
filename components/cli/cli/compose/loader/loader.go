package loader

import (
	"fmt"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"strings"

	"github.com/Sirupsen/logrus"
	"github.com/docker/cli/cli/compose/interpolation"
	"github.com/docker/cli/cli/compose/schema"
	"github.com/docker/cli/cli/compose/template"
	"github.com/docker/cli/cli/compose/types"
	"github.com/docker/cli/opts"
	"github.com/docker/go-connections/nat"
	units "github.com/docker/go-units"
	shellwords "github.com/mattn/go-shellwords"
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
)

// ParseYAML reads the bytes from a file, parses the bytes into a mapping
// structure, and returns it.
func ParseYAML(source []byte) (map[string]interface{}, error) {
	var cfg interface{}
	if err := yaml.Unmarshal(source, &cfg); err != nil {
		return nil, err
	}
	cfgMap, ok := cfg.(map[interface{}]interface{})
	if !ok {
		return nil, errors.Errorf("Top-level object must be a mapping")
	}
	converted, err := convertToStringKeysRecursive(cfgMap, "")
	if err != nil {
		return nil, err
	}
	return converted.(map[string]interface{}), nil
}

// Load reads a ConfigDetails and returns a fully loaded configuration
func Load(configDetails types.ConfigDetails) (*types.Config, error) {
	if len(configDetails.ConfigFiles) < 1 {
		return nil, errors.Errorf("No files specified")
	}
	if len(configDetails.ConfigFiles) > 1 {
		return nil, errors.Errorf("Multiple files are not yet supported")
	}

	configDict := getConfigDict(configDetails)

	if services, ok := configDict["services"]; ok {
		if servicesDict, ok := services.(map[string]interface{}); ok {
			forbidden := getProperties(servicesDict, types.ForbiddenProperties)

			if len(forbidden) > 0 {
				return nil, &ForbiddenPropertiesError{Properties: forbidden}
			}
		}
	}

	if err := schema.Validate(configDict, schema.Version(configDict)); err != nil {
		return nil, err
	}

	cfg := types.Config{}

	config, err := interpolateConfig(configDict, configDetails.LookupEnv)
	if err != nil {
		return nil, err
	}

	cfg.Services, err = LoadServices(config["services"], configDetails.WorkingDir, configDetails.LookupEnv)
	if err != nil {
		return nil, err
	}

	cfg.Networks, err = LoadNetworks(config["networks"])
	if err != nil {
		return nil, err
	}

	cfg.Volumes, err = LoadVolumes(config["volumes"])
	if err != nil {
		return nil, err
	}

	cfg.Secrets, err = LoadSecrets(config["secrets"], configDetails.WorkingDir)
	if err != nil {
		return nil, err
	}

	cfg.Configs, err = LoadConfigObjs(config["configs"], configDetails.WorkingDir)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func interpolateConfig(configDict map[string]interface{}, lookupEnv template.Mapping) (map[string]map[string]interface{}, error) {
	config := make(map[string]map[string]interface{})

	for _, key := range []string{"services", "networks", "volumes", "secrets", "configs"} {
		section, ok := configDict[key]
		if !ok {
			config[key] = make(map[string]interface{})
			continue
		}
		var err error
		config[key], err = interpolation.Interpolate(section.(map[string]interface{}), key, lookupEnv)
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

// GetUnsupportedProperties returns the list of any unsupported properties that are
// used in the Compose files.
func GetUnsupportedProperties(configDetails types.ConfigDetails) []string {
	unsupported := map[string]bool{}

	for _, service := range getServices(getConfigDict(configDetails)) {
		serviceDict := service.(map[string]interface{})
		for _, property := range types.UnsupportedProperties {
			if _, isSet := serviceDict[property]; isSet {
				unsupported[property] = true
			}
		}
	}

	return sortedKeys(unsupported)
}

func sortedKeys(set map[string]bool) []string {
	var keys []string
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}

// GetDeprecatedProperties returns the list of any deprecated properties that
// are used in the compose files.
func GetDeprecatedProperties(configDetails types.ConfigDetails) map[string]string {
	return getProperties(getServices(getConfigDict(configDetails)), types.DeprecatedProperties)
}

func getProperties(services map[string]interface{}, propertyMap map[string]string) map[string]string {
	output := map[string]string{}

	for _, service := range services {
		if serviceDict, ok := service.(map[string]interface{}); ok {
			for property, description := range propertyMap {
				if _, isSet := serviceDict[property]; isSet {
					output[property] = description
				}
			}
		}
	}

	return output
}

// ForbiddenPropertiesError is returned when there are properties in the Compose
// file that are forbidden.
type ForbiddenPropertiesError struct {
	Properties map[string]string
}

func (e *ForbiddenPropertiesError) Error() string {
	return "Configuration contains forbidden properties"
}

// TODO: resolve multiple files into a single config
func getConfigDict(configDetails types.ConfigDetails) map[string]interface{} {
	return configDetails.ConfigFiles[0].Config
}

func getServices(configDict map[string]interface{}) map[string]interface{} {
	if services, ok := configDict["services"]; ok {
		if servicesDict, ok := services.(map[string]interface{}); ok {
			return servicesDict
		}
	}

	return map[string]interface{}{}
}

func transform(source map[string]interface{}, target interface{}) error {
	data := mapstructure.Metadata{}
	config := &mapstructure.DecoderConfig{
		DecodeHook: mapstructure.ComposeDecodeHookFunc(
			createTransformHook(),
			mapstructure.StringToTimeDurationHookFunc()),
		Result:   target,
		Metadata: &data,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	return decoder.Decode(source)
}

func createTransformHook() mapstructure.DecodeHookFuncType {
	transforms := map[reflect.Type]func(interface{}) (interface{}, error){
		reflect.TypeOf(types.External{}):                         transformExternal,
		reflect.TypeOf(types.HealthCheckTest{}):                  transformHealthCheckTest,
		reflect.TypeOf(types.ShellCommand{}):                     transformShellCommand,
		reflect.TypeOf(types.StringList{}):                       transformStringList,
		reflect.TypeOf(map[string]string{}):                      transformMapStringString,
		reflect.TypeOf(types.UlimitsConfig{}):                    transformUlimits,
		reflect.TypeOf(types.UnitBytes(0)):                       transformSize,
		reflect.TypeOf([]types.ServicePortConfig{}):              transformServicePort,
		reflect.TypeOf(types.ServiceSecretConfig{}):              transformStringSourceMap,
		reflect.TypeOf(types.ServiceConfigObjConfig{}):           transformStringSourceMap,
		reflect.TypeOf(types.StringOrNumberList{}):               transformStringOrNumberList,
		reflect.TypeOf(map[string]*types.ServiceNetworkConfig{}): transformServiceNetworkMap,
		reflect.TypeOf(types.MappingWithEquals{}):                transformMappingOrListFunc("=", true),
		reflect.TypeOf(types.Labels{}):                           transformMappingOrListFunc("=", false),
		reflect.TypeOf(types.MappingWithColon{}):                 transformMappingOrListFunc(":", false),
		reflect.TypeOf(types.ServiceVolumeConfig{}):              transformServiceVolumeConfig,
	}

	return func(_ reflect.Type, target reflect.Type, data interface{}) (interface{}, error) {
		transform, ok := transforms[target]
		if !ok {
			return data, nil
		}
		return transform(data)
	}
}

// keys needs to be converted to strings for jsonschema
func convertToStringKeysRecursive(value interface{}, keyPrefix string) (interface{}, error) {
	if mapping, ok := value.(map[interface{}]interface{}); ok {
		dict := make(map[string]interface{})
		for key, entry := range mapping {
			str, ok := key.(string)
			if !ok {
				return nil, formatInvalidKeyError(keyPrefix, key)
			}
			var newKeyPrefix string
			if keyPrefix == "" {
				newKeyPrefix = str
			} else {
				newKeyPrefix = fmt.Sprintf("%s.%s", keyPrefix, str)
			}
			convertedEntry, err := convertToStringKeysRecursive(entry, newKeyPrefix)
			if err != nil {
				return nil, err
			}
			dict[str] = convertedEntry
		}
		return dict, nil
	}
	if list, ok := value.([]interface{}); ok {
		var convertedList []interface{}
		for index, entry := range list {
			newKeyPrefix := fmt.Sprintf("%s[%d]", keyPrefix, index)
			convertedEntry, err := convertToStringKeysRecursive(entry, newKeyPrefix)
			if err != nil {
				return nil, err
			}
			convertedList = append(convertedList, convertedEntry)
		}
		return convertedList, nil
	}
	return value, nil
}

func formatInvalidKeyError(keyPrefix string, key interface{}) error {
	var location string
	if keyPrefix == "" {
		location = "at top level"
	} else {
		location = fmt.Sprintf("in %s", keyPrefix)
	}
	return errors.Errorf("Non-string key %s: %#v", location, key)
}

// LoadServices produces a ServiceConfig map from a compose file Dict
// the servicesDict is not validated if directly used. Use Load() to enable validation
func LoadServices(servicesDict map[string]interface{}, workingDir string, lookupEnv template.Mapping) ([]types.ServiceConfig, error) {
	var services []types.ServiceConfig

	for name, serviceDef := range servicesDict {
		serviceConfig, err := LoadService(name, serviceDef.(map[string]interface{}), workingDir, lookupEnv)
		if err != nil {
			return nil, err
		}
		services = append(services, *serviceConfig)
	}

	return services, nil
}

// LoadService produces a single ServiceConfig from a compose file Dict
// the serviceDict is not validated if directly used. Use Load() to enable validation
func LoadService(name string, serviceDict map[string]interface{}, workingDir string, lookupEnv template.Mapping) (*types.ServiceConfig, error) {
	serviceConfig := &types.ServiceConfig{}
	if err := transform(serviceDict, serviceConfig); err != nil {
		return nil, err
	}
	serviceConfig.Name = name

	if err := resolveEnvironment(serviceConfig, workingDir, lookupEnv); err != nil {
		return nil, err
	}

	resolveVolumePaths(serviceConfig.Volumes, workingDir, lookupEnv)
	return serviceConfig, nil
}

func updateEnvironment(environment map[string]*string, vars map[string]*string, lookupEnv template.Mapping) {
	for k, v := range vars {
		interpolatedV, ok := lookupEnv(k)
		if (v == nil || *v == "") && ok {
			// lookupEnv is prioritized over vars
			environment[k] = &interpolatedV
		} else {
			environment[k] = v
		}
	}
}

func resolveEnvironment(serviceConfig *types.ServiceConfig, workingDir string, lookupEnv template.Mapping) error {
	environment := make(map[string]*string)

	if len(serviceConfig.EnvFile) > 0 {
		var envVars []string

		for _, file := range serviceConfig.EnvFile {
			filePath := absPath(workingDir, file)
			fileVars, err := opts.ParseEnvFile(filePath)
			if err != nil {
				return err
			}
			envVars = append(envVars, fileVars...)
		}
		updateEnvironment(environment,
			opts.ConvertKVStringsToMapWithNil(envVars), lookupEnv)
	}

	updateEnvironment(environment, serviceConfig.Environment, lookupEnv)
	serviceConfig.Environment = environment
	return nil
}

func resolveVolumePaths(volumes []types.ServiceVolumeConfig, workingDir string, lookupEnv template.Mapping) {
	for i, volume := range volumes {
		if volume.Type != "bind" {
			continue
		}

		filePath := expandUser(volume.Source, lookupEnv)
		// Check for a Unix absolute path first, to handle a Windows client
		// with a Unix daemon. This handles a Windows client connecting to a
		// Unix daemon. Note that this is not required for Docker for Windows
		// when specifying a local Windows path, because Docker for Windows
		// translates the Windows path into a valid path within the VM.
		if !path.IsAbs(filePath) {
			filePath = absPath(workingDir, filePath)
		}
		volume.Source = filePath
		volumes[i] = volume
	}
}

// TODO: make this more robust
func expandUser(path string, lookupEnv template.Mapping) string {
	if strings.HasPrefix(path, "~") {
		home, ok := lookupEnv("HOME")
		if !ok {
			logrus.Warn("cannot expand '~', because the environment lacks HOME")
			return path
		}
		return strings.Replace(path, "~", home, 1)
	}
	return path
}

func transformUlimits(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case int:
		return types.UlimitsConfig{Single: value}, nil
	case map[string]interface{}:
		ulimit := types.UlimitsConfig{}
		ulimit.Soft = value["soft"].(int)
		ulimit.Hard = value["hard"].(int)
		return ulimit, nil
	default:
		return data, errors.Errorf("invalid type %T for ulimits", value)
	}
}

// LoadNetworks produces a NetworkConfig map from a compose file Dict
// the source Dict is not validated if directly used. Use Load() to enable validation
func LoadNetworks(source map[string]interface{}) (map[string]types.NetworkConfig, error) {
	networks := make(map[string]types.NetworkConfig)
	err := transform(source, &networks)
	if err != nil {
		return networks, err
	}
	for name, network := range networks {
		if network.External.External && network.External.Name == "" {
			network.External.Name = name
			networks[name] = network
		}
	}
	return networks, nil
}

func externalVolumeError(volume, key string) error {
	return errors.Errorf(
		"conflicting parameters \"external\" and %q specified for volume %q",
		key, volume)
}

// LoadVolumes produces a VolumeConfig map from a compose file Dict
// the source Dict is not validated if directly used. Use Load() to enable validation
func LoadVolumes(source map[string]interface{}) (map[string]types.VolumeConfig, error) {
	volumes := make(map[string]types.VolumeConfig)
	err := transform(source, &volumes)
	if err != nil {
		return volumes, err
	}
	for name, volume := range volumes {
		if volume.External.External {
			if volume.Driver != "" {
				return nil, externalVolumeError(name, "driver")
			}
			if len(volume.DriverOpts) > 0 {
				return nil, externalVolumeError(name, "driver_opts")
			}
			if len(volume.Labels) > 0 {
				return nil, externalVolumeError(name, "labels")
			}
			if volume.External.Name == "" {
				volume.External.Name = name
				volumes[name] = volume
			} else {
				logrus.Warnf("volume %s: volume.external.name is deprecated in favor of volume.name", name)

				if volume.Name != "" {
					return nil, errors.Errorf("volume %s: volume.external.name and volume.name conflict; only use volume.name", name)
				}
			}
		}
	}
	return volumes, nil
}

// LoadSecrets produces a SecretConfig map from a compose file Dict
// the source Dict is not validated if directly used. Use Load() to enable validation
func LoadSecrets(source map[string]interface{}, workingDir string) (map[string]types.SecretConfig, error) {
	secrets := make(map[string]types.SecretConfig)
	if err := transform(source, &secrets); err != nil {
		return secrets, err
	}
	for name, secret := range secrets {
		if secret.External.External && secret.External.Name == "" {
			secret.External.Name = name
			secrets[name] = secret
		}
		if secret.File != "" {
			secret.File = absPath(workingDir, secret.File)
		}
	}
	return secrets, nil
}

// LoadConfigObjs produces a ConfigObjConfig map from a compose file Dict
// the source Dict is not validated if directly used. Use Load() to enable validation
func LoadConfigObjs(source map[string]interface{}, workingDir string) (map[string]types.ConfigObjConfig, error) {
	configs := make(map[string]types.ConfigObjConfig)
	if err := transform(source, &configs); err != nil {
		return configs, err
	}
	for name, config := range configs {
		if config.External.External && config.External.Name == "" {
			config.External.Name = name
			configs[name] = config
		}
		if config.File != "" {
			config.File = absPath(workingDir, config.File)
		}
	}
	return configs, nil
}

func absPath(workingDir string, filePath string) string {
	if filepath.IsAbs(filePath) {
		return filePath
	}
	return filepath.Join(workingDir, filePath)
}

func transformMapStringString(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case map[string]interface{}:
		return toMapStringString(value, false), nil
	case map[string]string:
		return value, nil
	default:
		return data, errors.Errorf("invalid type %T for map[string]string", value)
	}
}

func transformExternal(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case bool:
		return map[string]interface{}{"external": value}, nil
	case map[string]interface{}:
		return map[string]interface{}{"external": true, "name": value["name"]}, nil
	default:
		return data, errors.Errorf("invalid type %T for external", value)
	}
}

func transformServicePort(data interface{}) (interface{}, error) {
	switch entries := data.(type) {
	case []interface{}:
		// We process the list instead of individual items here.
		// The reason is that one entry might be mapped to multiple ServicePortConfig.
		// Therefore we take an input of a list and return an output of a list.
		ports := []interface{}{}
		for _, entry := range entries {
			switch value := entry.(type) {
			case int:
				v, err := toServicePortConfigs(fmt.Sprint(value))
				if err != nil {
					return data, err
				}
				ports = append(ports, v...)
			case string:
				v, err := toServicePortConfigs(value)
				if err != nil {
					return data, err
				}
				ports = append(ports, v...)
			case map[string]interface{}:
				ports = append(ports, value)
			default:
				return data, errors.Errorf("invalid type %T for port", value)
			}
		}
		return ports, nil
	default:
		return data, errors.Errorf("invalid type %T for port", entries)
	}
}

func transformStringSourceMap(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case string:
		return map[string]interface{}{"source": value}, nil
	case map[string]interface{}:
		return data, nil
	default:
		return data, errors.Errorf("invalid type %T for secret", value)
	}
}

func transformServiceVolumeConfig(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case string:
		return ParseVolume(value)
	case map[string]interface{}:
		return data, nil
	default:
		return data, errors.Errorf("invalid type %T for service volume", value)
	}

}

func transformServiceNetworkMap(value interface{}) (interface{}, error) {
	if list, ok := value.([]interface{}); ok {
		mapValue := map[interface{}]interface{}{}
		for _, name := range list {
			mapValue[name] = nil
		}
		return mapValue, nil
	}
	return value, nil
}

func transformStringOrNumberList(value interface{}) (interface{}, error) {
	list := value.([]interface{})
	result := make([]string, len(list))
	for i, item := range list {
		result[i] = fmt.Sprint(item)
	}
	return result, nil
}

func transformStringList(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case string:
		return []string{value}, nil
	case []interface{}:
		return value, nil
	default:
		return data, errors.Errorf("invalid type %T for string list", value)
	}
}

func transformMappingOrListFunc(sep string, allowNil bool) func(interface{}) (interface{}, error) {
	return func(data interface{}) (interface{}, error) {
		return transformMappingOrList(data, sep, allowNil), nil
	}
}

func transformMappingOrList(mappingOrList interface{}, sep string, allowNil bool) interface{} {
	switch value := mappingOrList.(type) {
	case map[string]interface{}:
		return toMapStringString(value, allowNil)
	case ([]interface{}):
		result := make(map[string]interface{})
		for _, value := range value {
			parts := strings.SplitN(value.(string), sep, 2)
			key := parts[0]
			switch {
			case len(parts) == 1 && allowNil:
				result[key] = nil
			case len(parts) == 1 && !allowNil:
				result[key] = ""
			default:
				result[key] = parts[1]
			}
		}
		return result
	}
	panic(errors.Errorf("expected a map or a list, got %T: %#v", mappingOrList, mappingOrList))
}

func transformShellCommand(value interface{}) (interface{}, error) {
	if str, ok := value.(string); ok {
		return shellwords.Parse(str)
	}
	return value, nil
}

func transformHealthCheckTest(data interface{}) (interface{}, error) {
	switch value := data.(type) {
	case string:
		return append([]string{"CMD-SHELL"}, value), nil
	case []interface{}:
		return value, nil
	default:
		return value, errors.Errorf("invalid type %T for healthcheck.test", value)
	}
}

func transformSize(value interface{}) (interface{}, error) {
	switch value := value.(type) {
	case int:
		return int64(value), nil
	case string:
		return units.RAMInBytes(value)
	}
	panic(errors.Errorf("invalid type for size %T", value))
}

func toServicePortConfigs(value string) ([]interface{}, error) {
	var portConfigs []interface{}

	ports, portBindings, err := nat.ParsePortSpecs([]string{value})
	if err != nil {
		return nil, err
	}
	// We need to sort the key of the ports to make sure it is consistent
	keys := []string{}
	for port := range ports {
		keys = append(keys, string(port))
	}
	sort.Strings(keys)

	for _, key := range keys {
		// Reuse ConvertPortToPortConfig so that it is consistent
		portConfig, err := opts.ConvertPortToPortConfig(nat.Port(key), portBindings)
		if err != nil {
			return nil, err
		}
		for _, p := range portConfig {
			portConfigs = append(portConfigs, types.ServicePortConfig{
				Protocol:  string(p.Protocol),
				Target:    p.TargetPort,
				Published: p.PublishedPort,
				Mode:      string(p.PublishMode),
			})
		}
	}

	return portConfigs, nil
}

func toMapStringString(value map[string]interface{}, allowNil bool) map[string]interface{} {
	output := make(map[string]interface{})
	for key, value := range value {
		output[key] = toString(value, allowNil)
	}
	return output
}

func toString(value interface{}, allowNil bool) interface{} {
	switch {
	case value != nil:
		return fmt.Sprint(value)
	case allowNil:
		return nil
	default:
		return ""
	}
}
