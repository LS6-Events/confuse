package confuse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParseMapToStruct(t *testing.T) {
	t.Run("should parse map to struct", func(t *testing.T) {
		var s struct {
			Hello string
			World string
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": "World",
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, "World", s.World)
	})

	t.Run("should parse map to struct with nested struct", func(t *testing.T) {
		var s struct {
			Hello string
			World struct {
				Foo string
				Bar int
			}
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": map[string]any{
				"Foo": "Foo",
				"Bar": 42,
			},
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, "Foo", s.World.Foo)
		require.Equal(t, 42, s.World.Bar)
	})

	t.Run("should parse map to struct with nested struct and slice", func(t *testing.T) {
		var s struct {
			Hello string
			World struct {
				Foo string
				Bar []int
			}
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": map[string]any{
				"Foo": "Foo",
				"Bar": []int{1, 2, 3},
			},
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, "Foo", s.World.Foo)
		require.Equal(t, []int{1, 2, 3}, s.World.Bar)
	})

	t.Run("should error on invalid type", func(t *testing.T) {
		var s struct {
			Hello string
			World float64
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": "World",
		}

		err := parseMapToStruct(&s, m)
		require.Error(t, err)
	})

	t.Run("should error on invalid type in nested struct", func(t *testing.T) {
		var s struct {
			Hello string
			World struct {
				Foo string
				Bar int
			}
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": map[string]any{
				"Foo": "Foo",
				"Bar": "Bar",
			},
		}

		err := parseMapToStruct(&s, m)
		require.Error(t, err)
	})

	t.Run("should error on invalid type in nested struct and slice", func(t *testing.T) {
		var s struct {
			Hello string
			World struct {
				Foo string
				Bar []int
			}
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": map[string]any{
				"Foo": "Foo",
				"Bar": []any{"Bar"},
			},
		}

		err := parseMapToStruct(&s, m)
		require.Error(t, err)
	})

	t.Run("should be able to convert between int and float", func(t *testing.T) {
		var s struct {
			Hello string
			World float64
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": 42,
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, float64(42), s.World)
	})

	t.Run("should be able to convert between string and int", func(t *testing.T) {
		var s struct {
			Hello string
			World int
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": "42",
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, 42, s.World)
	})

	t.Run("should be able to convert between string and bool", func(t *testing.T) {
		var s struct {
			Hello string
			World bool
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": "true",
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, true, s.World)
	})

	t.Run("should be able to use the config tag", func(t *testing.T) {
		var s struct {
			Hello string `config:"foo"`
			World string `config:"bar"`
		}

		m := map[string]any{
			"foo": "Hello",
			"bar": "World",
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, "World", s.World)
	})

	t.Run("should be able to parse slices, maps and arrays", func(t *testing.T) {
		var s struct {
			Slice []int
			Map   map[string]int
			Array [3]int
		}

		m := map[string]any{
			"Slice": []any{1, 2, 3},
			"Map": map[string]any{
				"foo": 1,
				"bar": 2,
			},
			"Array": []any{1, 2, 3},
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)
		require.Equal(t, []int{1, 2, 3}, s.Slice)
		require.Equal(t, map[string]int{"foo": 1, "bar": 2}, s.Map)
		require.Equal(t, [3]int{1, 2, 3}, s.Array)
	})

	t.Run("should be able to parse types with their values being strings", func(t *testing.T) {
		var s struct {
			Hello string
			World int
			Test  bool
			Float float64
		}

		m := map[string]any{
			"Hello": "Hello",
			"World": "42",
			"Test":  "true",
			"Float": "42.42",
		}

		err := parseMapToStruct(&s, m)
		require.NoError(t, err)

		require.Equal(t, "Hello", s.Hello)
		require.Equal(t, 42, s.World)
		require.Equal(t, true, s.Test)
		require.Equal(t, 42.42, s.Float)
	})
}
