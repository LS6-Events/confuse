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

	require.True(t, config.Env.Env)
	require.True(t, config.Json.Json)
	require.True(t, config.Toml.Toml)
	require.True(t, config.Yaml.Yaml)
}
