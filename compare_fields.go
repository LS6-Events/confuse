package confuse

import "reflect"

// findIndexOfField compares a field to a list of strings.
// it returns the index of the first match, or -1 if no match is found.
func findIndexOfField(field reflect.StructField, toCompare []string) int {
	for i, compare := range toCompare {
		if compareField(field, compare) {
			return i
		}
	}

	return -1
}
