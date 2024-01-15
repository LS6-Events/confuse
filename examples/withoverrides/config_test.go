package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	config, err := getConfig()
	require.NoError(t, err)

	require.Equal(t, "bar", config.Foo)
	require.Equal(t, 42, config.Bar)

	require.Equal(t, "localhost2", config.DatabaseConfig.Host)
	require.Equal(t, 8083, config.DatabaseConfig.Port)
}
