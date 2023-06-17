//go:build integration
// +build integration

package dns

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_integration(t *testing.T) {
	t.Parallel()

	fetcher, err := New(SetProviders(Cloudflare, OpenDNS))
	require.NoError(t, err)

	ctx := context.Background()

	publicIP1, err := fetcher.IP4(ctx)
	require.NoError(t, err)
	assert.NotNil(t, publicIP1)

	publicIP2, err := fetcher.IP4(ctx)
	require.NoError(t, err)
	assert.NotNil(t, publicIP2)

	assert.Equal(t, publicIP1.String(), publicIP2.String())

	t.Logf("Public IP is %s", publicIP1)
}
