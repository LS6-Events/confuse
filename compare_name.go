package confuse

import (
	"github.com/iancoleman/strcase"
)

func compareName(mapKey, fieldName string) bool {
	comparisonValue := fieldName

	// First check if the map key matches the field name
	if mapKey == comparisonValue {
		return true
	}

	// Next check if the field name converted to snake case matches the map key
	comparisonValue = strcase.ToSnake(fieldName)
	if mapKey == comparisonValue {
		return true
	}

	// Next check if the field name converted to screaming snake case matches the toCompare string
	comparisonValue = strcase.ToScreamingSnake(fieldName)
	if mapKey == comparisonValue {
		return true
	}

	// Next check if the field name converted to kebab case matches the toCompare string
	comparisonValue = strcase.ToKebab(fieldName)
	if mapKey == comparisonValue {
		return true
	}

	// Next check if the field name converted to camel case matches the toCompare string
	comparisonValue = strcase.ToCamel(fieldName)
	if mapKey == comparisonValue {
		return true
	}

	// Next check if the field name converted to lower camel case matches the toCompare string
	comparisonValue = strcase.ToLowerCamel(fieldName)
	if mapKey == comparisonValue {
		return true
	}

	return false
}
