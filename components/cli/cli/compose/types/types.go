package types

import (
	"time"
)

// UnsupportedProperties not yet supported by this implementation of the compose file
var UnsupportedProperties = []string{
	"build",
	"cap_add",
	"cap_drop",
	"cgroup_parent",
	"devices",
	"domainname",
	"external_links",
	"ipc",
	"links",
	"mac_address",
	"network_mode",
	"privileged",
	"restart",
	"security_opt",
	"shm_size",
	"sysctls",
	"tmpfs",
	"userns_mode",
}

// DeprecatedProperties that were removed from the v3 format, but their
// use should not impact the behaviour of the application.
var DeprecatedProperties = map[string]string{
	"container_name": "Setting the container name is not supported.",
	"expose":         "Exposing ports is unnecessary - services on the same network can access each other's containers on any port.",
}

// ForbiddenProperties that are not supported in this implementation of the
// compose file.
var ForbiddenProperties = map[string]string{
	"extends":       "Support for `extends` is not implemented yet.",
	"volume_driver": "Instead of setting the volume driver on the service, define a volume using the top-level `volumes` option and specify the driver there.",
	"volumes_from":  "To share a volume between services, define it using the top-level `volumes` option and reference it from each service that shares it using the service-level `volumes` option.",
	"cpu_quota":     "Set resource limits using deploy.resources",
	"cpu_shares":    "Set resource limits using deploy.resources",
	"cpuset":        "Set resource limits using deploy.resources",
	"mem_limit":     "Set resource limits using deploy.resources",
	"memswap_limit": "Set resource limits using deploy.resources",
}

// ConfigFile is a filename and the contents of the file as a Dict
type ConfigFile struct {
	Filename string
	Config   map[string]interface{}
}

// ConfigDetails are the details about a group of ConfigFiles
type ConfigDetails struct {
	WorkingDir  string
	ConfigFiles []ConfigFile
	Environment map[string]string
}

// LookupEnv provides a lookup function for environment variables
func (cd ConfigDetails) LookupEnv(key string) (string, bool) {
	v, ok := cd.Environment[key]
	return v, ok
}

// Config is a full compose file configuration
type Config struct {
	Services []ServiceConfig
	Networks map[string]NetworkConfig
	Volumes  map[string]VolumeConfig
	Secrets  map[string]SecretConfig
	Configs  map[string]ConfigObjConfig
}

// ServiceConfig is the configuration of one service
type ServiceConfig struct {
	Name string

	CapAdd          []string `mapstructure:"cap_add"`
	CapDrop         []string `mapstructure:"cap_drop"`
	CgroupParent    string   `mapstructure:"cgroup_parent"`
	Command         ShellCommand
	Configs         []ServiceConfigObjConfig
	ContainerName   string               `mapstructure:"container_name"`
	CredentialSpec  CredentialSpecConfig `mapstructure:"credential_spec"`
	DependsOn       []string             `mapstructure:"depends_on"`
	Deploy          DeployConfig
	Devices         []string
	DNS             StringList
	DNSSearch       StringList `mapstructure:"dns_search"`
	DomainName      string     `mapstructure:"domainname"`
	Entrypoint      ShellCommand
	Environment     MappingWithEquals
	EnvFile         StringList `mapstructure:"env_file"`
	Expose          StringOrNumberList
	ExternalLinks   []string         `mapstructure:"external_links"`
	ExtraHosts      MappingWithColon `mapstructure:"extra_hosts"`
	Hostname        string
	HealthCheck     *HealthCheckConfig
	Image           string
	Ipc             string
	Labels          Labels
	Links           []string
	Logging         *LoggingConfig
	MacAddress      string `mapstructure:"mac_address"`
	NetworkMode     string `mapstructure:"network_mode"`
	Networks        map[string]*ServiceNetworkConfig
	Pid             string
	Ports           []ServicePortConfig
	Privileged      bool
	ReadOnly        bool `mapstructure:"read_only"`
	Restart         string
	Secrets         []ServiceSecretConfig
	SecurityOpt     []string       `mapstructure:"security_opt"`
	StdinOpen       bool           `mapstructure:"stdin_open"`
	StopGracePeriod *time.Duration `mapstructure:"stop_grace_period"`
	StopSignal      string         `mapstructure:"stop_signal"`
	Tmpfs           StringList
	Tty             bool `mapstructure:"tty"`
	Ulimits         map[string]*UlimitsConfig
	User            string
	Volumes         []ServiceVolumeConfig
	WorkingDir      string `mapstructure:"working_dir"`
}

// ShellCommand is a string or list of string args
type ShellCommand []string

// StringList is a type for fields that can be a string or list of strings
type StringList []string

// StringOrNumberList is a type for fields that can be a list of strings or
// numbers
type StringOrNumberList []string

// MappingWithEquals is a mapping type that can be converted from a list of
// key[=value] strings.
// For the key with an empty value (`key=`), the mapped value is set to a pointer to `""`.
// For the key without value (`key`), the mapped value is set to nil.
type MappingWithEquals map[string]*string

// Labels is a mapping type for labels
type Labels map[string]string

// MappingWithColon is a mapping type that can be converted from a list of
// 'key: value' strings
type MappingWithColon map[string]string

// LoggingConfig the logging configuration for a service
type LoggingConfig struct {
	Driver  string
	Options map[string]string
}

// DeployConfig the deployment configuration for a service
type DeployConfig struct {
	Mode          string
	Replicas      *uint64
	Labels        Labels
	UpdateConfig  *UpdateConfig `mapstructure:"update_config"`
	Resources     Resources
	RestartPolicy *RestartPolicy `mapstructure:"restart_policy"`
	Placement     Placement
	EndpointMode  string `mapstructure:"endpoint_mode"`
}

// HealthCheckConfig the healthcheck configuration for a service
type HealthCheckConfig struct {
	Test        HealthCheckTest
	Timeout     string
	Interval    string
	Retries     *uint64
	StartPeriod string
	Disable     bool
}

// HealthCheckTest is the command run to test the health of a service
type HealthCheckTest []string

// UpdateConfig the service update configuration
type UpdateConfig struct {
	Parallelism     *uint64
	Delay           time.Duration
	FailureAction   string `mapstructure:"failure_action"`
	Monitor         time.Duration
	MaxFailureRatio float32 `mapstructure:"max_failure_ratio"`
	Order           string
}

// Resources the resource limits and reservations
type Resources struct {
	Limits       *Resource
	Reservations *Resource
}

// Resource is a resource to be limited or reserved
type Resource struct {
	// TODO: types to convert from units and ratios
	NanoCPUs    string    `mapstructure:"cpus"`
	MemoryBytes UnitBytes `mapstructure:"memory"`
}

// UnitBytes is the bytes type
type UnitBytes int64

// RestartPolicy the service restart policy
type RestartPolicy struct {
	Condition   string
	Delay       *time.Duration
	MaxAttempts *uint64 `mapstructure:"max_attempts"`
	Window      *time.Duration
}

// Placement constraints for the service
type Placement struct {
	Constraints []string
	Preferences []PlacementPreferences
}

// PlacementPreferences is the preferences for a service placement
type PlacementPreferences struct {
	Spread string
}

// ServiceNetworkConfig is the network configuration for a service
type ServiceNetworkConfig struct {
	Aliases     []string
	Ipv4Address string `mapstructure:"ipv4_address"`
	Ipv6Address string `mapstructure:"ipv6_address"`
}

// ServicePortConfig is the port configuration for a service
type ServicePortConfig struct {
	Mode      string
	Target    uint32
	Published uint32
	Protocol  string
}

// ServiceVolumeConfig are references to a volume used by a service
type ServiceVolumeConfig struct {
	Type        string
	Source      string
	Target      string
	ReadOnly    bool `mapstructure:"read_only"`
	Consistency string
	Bind        *ServiceVolumeBind
	Volume      *ServiceVolumeVolume
}

// ServiceVolumeBind are options for a service volume of type bind
type ServiceVolumeBind struct {
	Propagation string
}

// ServiceVolumeVolume are options for a service volume of type volume
type ServiceVolumeVolume struct {
	NoCopy bool `mapstructure:"nocopy"`
}

type fileReferenceConfig struct {
	Source string
	Target string
	UID    string
	GID    string
	Mode   *uint32
}

// ServiceConfigObjConfig is the config obj configuration for a service
type ServiceConfigObjConfig fileReferenceConfig

// ServiceSecretConfig is the secret configuration for a service
type ServiceSecretConfig fileReferenceConfig

// UlimitsConfig the ulimit configuration
type UlimitsConfig struct {
	Single int
	Soft   int
	Hard   int
}

// NetworkConfig for a network
type NetworkConfig struct {
	Driver     string
	DriverOpts map[string]string `mapstructure:"driver_opts"`
	Ipam       IPAMConfig
	External   External
	Internal   bool
	Attachable bool
	Labels     Labels
}

// IPAMConfig for a network
type IPAMConfig struct {
	Driver string
	Config []*IPAMPool
}

// IPAMPool for a network
type IPAMPool struct {
	Subnet string
}

// VolumeConfig for a volume
type VolumeConfig struct {
	Name       string
	Driver     string
	DriverOpts map[string]string `mapstructure:"driver_opts"`
	External   External
	Labels     Labels
}

// External identifies a Volume or Network as a reference to a resource that is
// not managed, and should already exist.
// External.name is deprecated and replaced by Volume.name
type External struct {
	Name     string
	External bool
}

// CredentialSpecConfig for credential spec on Windows
type CredentialSpecConfig struct {
	File     string
	Registry string
}

type fileObjectConfig struct {
	File     string
	External External
	Labels   Labels
}

// SecretConfig for a secret
type SecretConfig fileObjectConfig

// ConfigObjConfig is the config for the swarm "Config" object
type ConfigObjConfig fileObjectConfig
