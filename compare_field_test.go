package confuse

import (
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestCompareField(t *testing.T) {
	var s struct {
		Foo             string `config:"foo"`
		TwoCases        string
		privateTwoCases string
	}

	rS := reflect.TypeOf(s)

	t.Run("compares field with config struct tag", func(t *testing.T) {
		result := compareField(rS.Field(0), "foo")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name", func(t *testing.T) {
		result := compareField(rS.Field(1), "TwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to snake case", func(t *testing.T) {
		result := compareField(rS.Field(1), "two_cases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to screaming snake case", func(t *testing.T) {
		result := compareField(rS.Field(1), "TWO_CASES")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to kebab case", func(t *testing.T) {
		result := compareField(rS.Field(1), "two-cases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to camel case", func(t *testing.T) {
		result := compareField(rS.Field(2), "PrivateTwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to lower camel case", func(t *testing.T) {
		result := compareField(rS.Field(1), "twoCases")
		require.Equal(t, true, result)
	})

	t.Run("returns false if no match", func(t *testing.T) {
		result := compareField(rS.Field(1), "no_match")
		require.Equal(t, false, result)
	})
}
