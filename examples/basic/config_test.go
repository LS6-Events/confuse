package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	config, err := getConfig()
	require.NoError(t, err)

	require.Equal(t, "bar", config.Foo)
	require.Equal(t, 42, config.Bar)

	require.Equal(t, "localhost", config.DatabaseConfig.Host)
	require.Equal(t, 8080, config.DatabaseConfig.Port)

	require.Equal(t, []int{1, 2, 3}, config.Slice)
	require.Equal(t, map[string]int{"test": 1, "test2": 2}, config.Map)
	require.Equal(t, [3]int{1, 2, 3}, config.Array)

	require.Equal(t, "world", config.DiffName.Hello)

	require.Equal(t, "world", config.PointerStruct.Hello)

	require.Equal(t, "2019-01-01T00:00:00Z", config.Date.Format(time.RFC3339))
}
