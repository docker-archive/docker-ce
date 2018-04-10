package container // import "github.com/docker/docker/integration/container"

import (
	"context"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/integration/internal/container"
	"github.com/docker/docker/integration/internal/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPsFilter(t *testing.T) {
	defer setupTest(t)()
	client := request.NewAPIClient(t)
	ctx := context.Background()

	prev := container.Create(t, ctx, client, container.WithName("prev"))
	container.Create(t, ctx, client, container.WithName("top"))
	next := container.Create(t, ctx, client, container.WithName("next"))

	containerIDs := func(containers []types.Container) []string {
		entries := []string{}
		for _, container := range containers {
			entries = append(entries, container.ID)
		}
		return entries
	}

	f1 := filters.NewArgs()
	f1.Add("since", "top")
	q1, err := client.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: f1,
	})
	require.NoError(t, err)
	assert.Contains(t, containerIDs(q1), next)

	f2 := filters.NewArgs()
	f2.Add("before", "top")
	q2, err := client.ContainerList(ctx, types.ContainerListOptions{
		All:     true,
		Filters: f2,
	})
	require.NoError(t, err)
	assert.Contains(t, containerIDs(q2), prev)
}
