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

	t.Run("sets schema correctly", func(t *testing.T) {
		_, err := getConfig(false)
		require.NoError(t, err)

		stat, err := os.Stat("./schema.json")
		require.NoError(t, err)
		require.NotNil(t, stat)

		schema, err := gabs.ParseJSONFile("./schema.json")
		require.NoError(t, err)

		require.Equal(t, []interface{}{"foo", "bar"}, schema.Path("required").Data().([]interface{}))

		properties := schema.Path("properties")

		require.Equal(t, "integer", properties.Path("bar.type").Data().(string))
		require.Equal(t, float64(1), properties.Path("bar.allOf.0.minimum").Data().(float64))
		require.True(t, properties.Path("bar.allOf.0.exclusiveMinimum").Data().(bool))
		require.Equal(t, float64(10), properties.Path("bar.allOf.1.maximum").Data().(float64))
		require.True(t, properties.Path("bar.allOf.1.exclusiveMaximum").Data().(bool))

		require.Equal(t, "string", properties.Path("foo.type").Data().(string))
		require.Equal(t, []interface{}{"foo", "bar"}, properties.Path("foo.allOf.0.enum").Data().([]interface{}))
	})

	t.Run("with validation pass", func(t *testing.T) {
		conf, err := getConfig(false)
		require.NoError(t, err)

		require.Equal(t, "bar", conf.Foo)
		require.Equal(t, 5, conf.Bar)
	})

	t.Run("with validation fail", func(t *testing.T) {
		_, err := getConfig(true)
		require.Error(t, err)
	})
}
