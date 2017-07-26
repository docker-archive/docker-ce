package service

import (
	"strings"

	"golang.org/x/net/context"

	"github.com/docker/cli/cli"
	"github.com/docker/cli/cli/command"
	"github.com/docker/cli/cli/command/formatter"
	"github.com/docker/docker/api/types"
	apiclient "github.com/docker/docker/client"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

type inspectOptions struct {
	refs   []string
	format string
	pretty bool
}

func newInspectCommand(dockerCli *command.DockerCli) *cobra.Command {
	var opts inspectOptions

	cmd := &cobra.Command{
		Use:   "inspect [OPTIONS] SERVICE [SERVICE...]",
		Short: "Display detailed information on one or more services",
		Args:  cli.RequiresMinArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			opts.refs = args

			if opts.pretty && len(opts.format) > 0 {
				return errors.Errorf("--format is incompatible with human friendly format")
			}
			return runInspect(dockerCli, opts)
		},
	}

	flags := cmd.Flags()
	flags.StringVarP(&opts.format, "format", "f", "", "Format the output using the given Go template")
	flags.BoolVar(&opts.pretty, "pretty", false, "Print the information in a human friendly format")
	return cmd
}

func runInspect(dockerCli *command.DockerCli, opts inspectOptions) error {
	client := dockerCli.Client()
	ctx := context.Background()

	if opts.pretty {
		opts.format = "pretty"
	}

	getRef := func(ref string) (interface{}, []byte, error) {
		// Service inspect shows defaults values in empty fields.
		service, _, err := client.ServiceInspectWithRaw(ctx, ref, types.ServiceInspectOptions{InsertDefaults: true})
		if err == nil || !apiclient.IsErrServiceNotFound(err) {
			return service, nil, err
		}
		return nil, nil, errors.Errorf("Error: no such service: %s", ref)
	}

	getNetwork := func(ref string) (interface{}, []byte, error) {
		network, _, err := client.NetworkInspectWithRaw(ctx, ref, types.NetworkInspectOptions{Scope: "swarm"})
		if err == nil || !apiclient.IsErrNetworkNotFound(err) {
			return network, nil, err
		}
		return nil, nil, errors.Errorf("Error: no such network: %s", ref)
	}

	f := opts.format
	if len(f) == 0 {
		f = "raw"
		if len(dockerCli.ConfigFile().ServiceInspectFormat) > 0 {
			f = dockerCli.ConfigFile().ServiceInspectFormat
		}
	}

	// check if the user is trying to apply a template to the pretty format, which
	// is not supported
	if strings.HasPrefix(f, "pretty") && f != "pretty" {
		return errors.Errorf("Cannot supply extra formatting options to the pretty template")
	}

	serviceCtx := formatter.Context{
		Output: dockerCli.Out(),
		Format: formatter.NewServiceFormat(f),
	}

	if err := formatter.ServiceInspectWrite(serviceCtx, opts.refs, getRef, getNetwork); err != nil {
		return cli.StatusError{StatusCode: 1, Status: err.Error()}
	}
	return nil
}
