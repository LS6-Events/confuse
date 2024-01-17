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

	require.Equal(t, "localhost", config.DatabaseConfig.Host)
	require.Equal(t, 8080, config.DatabaseConfig.Port)
}
