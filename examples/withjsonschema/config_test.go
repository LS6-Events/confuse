package main

import (
	"github.com/Jeffail/gabs/v2"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestConfig(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}

	_, err := getConfig()
	require.NoError(t, err)

	stat, err := os.Stat("./schema.json")
	require.NoError(t, err)
	require.NotNil(t, stat)

	schema, err := gabs.ParseJSONFile("./schema.json")
	require.NoError(t, err)

	properties := schema.Path("properties")

	require.Equal(t, "array", properties.Path("array.type").Data().(string))
	require.Equal(t, float64(3), properties.Path("array.minItems").Data().(float64))
	require.Equal(t, float64(3), properties.Path("array.maxItems").Data().(float64))

	require.Equal(t, "integer", properties.Path("bar.type").Data().(string))

	require.Equal(t, "object", properties.Path("database_config.type").Data().(string))
	require.Equal(t, "string", properties.Path("database_config.properties.host.type").Data().(string))
	require.Equal(t, "integer", properties.Path("database_config.properties.port.type").Data().(string))

	require.Equal(t, "string", properties.Path("foo.type").Data().(string))

	require.Equal(t, "object", properties.Path("map.type").Data().(string))
	require.Equal(t, "integer", properties.Path("map.additionalProperties.type").Data().(string))

	require.Equal(t, "array", properties.Path("slice.type").Data().(string))
	require.Equal(t, "integer", properties.Path("slice.items.type").Data().(string))
}
