package confuse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCompareName(t *testing.T) {
	t.Run("compares directly", func(t *testing.T) {
		result := compareName("TwoCases", "TwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to snake case", func(t *testing.T) {
		result := compareName("two_cases", "TwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to screaming snake case", func(t *testing.T) {
		result := compareName("TWO_CASES", "TwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to kebab case", func(t *testing.T) {
		result := compareName("two-cases", "TwoCases")
		require.Equal(t, true, result)
	})

	t.Run("compares field with field name converted to camel case", func(t *testing.T) {
		result := compareName("privateTwoCases", "PrivateTwoCases")
		require.Equal(t, true, result)
	})

	t.Run("returns false if no match", func(t *testing.T) {
		result := compareName("no_match", "TwoCases")
		require.Equal(t, false, result)
	})
}
