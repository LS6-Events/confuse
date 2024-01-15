package confuse

import (
	"github.com/iancoleman/strcase"
	"reflect"
)

const configTag = "config"

func compareField(field reflect.StructField, toCompare string) bool {
	// Check if the field has a config tag and if it matches the toCompare string
	configTagValue := field.Tag.Get(configTag)
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name matches the toCompare string
	configTagValue = field.Name
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name converted to snake case matches the toCompare string
	configTagValue = strcase.ToSnake(field.Name)
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name converted to screaming snake case matches the toCompare string
	configTagValue = strcase.ToScreamingSnake(field.Name)
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name converted to kebab case matches the toCompare string
	configTagValue = strcase.ToKebab(field.Name)
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name converted to camel case matches the toCompare string
	configTagValue = strcase.ToCamel(field.Name)
	if configTagValue == toCompare {
		return true
	}

	// Next check if the field name converted to lower camel case matches the toCompare string
	configTagValue = strcase.ToLowerCamel(field.Name)
	if configTagValue == toCompare {
		return true
	}

	return false
}
