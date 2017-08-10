package network

import (
	"io/ioutil"
	"testing"

	"github.com/docker/cli/cli/internal/test"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/pkg/testutil"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
)

func TestNetworkConnectErrors(t *testing.T) {
	testCases := []struct {
		args               []string
		networkConnectFunc func(ctx context.Context, networkID, container string, config *network.EndpointSettings) error
		expectedError      string
	}{
		{
			expectedError: "requires exactly 2 argument(s)",
		},
		{
			args: []string{"toto", "titi"},
			networkConnectFunc: func(ctx context.Context, networkID, container string, config *network.EndpointSettings) error {
				return errors.Errorf("error connecting network")
			},
			expectedError: "error connecting network",
		},
	}

	for _, tc := range testCases {
		cmd := newConnectCommand(
			test.NewFakeCli(&fakeClient{
				networkConnectFunc: tc.networkConnectFunc,
			}),
		)
		cmd.SetArgs(tc.args)
		cmd.SetOutput(ioutil.Discard)
		testutil.ErrorContains(t, cmd.Execute(), tc.expectedError)

	}
}

func TestNetworkConnectWithFlags(t *testing.T) {
	expectedOpts := []network.IPAMConfig{
		{
			"192.168.4.0/24",
			"192.168.4.0/24",
			"192.168.4.1/24",
			map[string]string{},
		},
	}
	cli := test.NewFakeCli(&fakeClient{
		networkConnectFunc: func(ctx context.Context, networkID, container string, config *network.EndpointSettings) error {
			assert.Equal(t, expectedOpts, config.IPAMConfig, "not expected driver error")
			return nil
		},
	})
	args := []string{"banana"}
	cmd := newCreateCommand(cli)

	cmd.SetArgs(args)
	cmd.Flags().Set("driver", "foo")
	cmd.Flags().Set("ip-range", "192.168.4.0/24")
	cmd.Flags().Set("gateway", "192.168.4.1/24")
	cmd.Flags().Set("subnet", "192.168.4.0/24")
	assert.NoError(t, cmd.Execute())
}
