package confuse

import (
	"errors"
	"reflect"

	"dario.cat/mergo"
	"github.com/ls6-events/validjsonator"
)

func (s *Service) schemaFromType(t reflect.Type, fuzzy bool) (validjsonator.Schema, error) {
	var schema validjsonator.Schema

	switch t.Kind() {
	case reflect.Struct:
		schema.Type = "object"
		schema.Properties = make(map[string]validjsonator.Schema)
		for i := 0; i < t.NumField(); i++ {
			field := t.Field(i)
			configTagName, _ := extractValuesFromTag(field.Tag.Get(configTag))

			if configTagName == "-" {
				continue
			}

			if configTagName == "" {
				configTagName = s.JSONSchemaKeyModifier(field.Name)
			}

			fieldSchema, err := s.schemaFromType(field.Type, fuzzy)
			if err != nil {
				return validjsonator.Schema{}, err
			}

			if s.ShouldValidate {
				validationSchema, required := validjsonator.ValidationTagsToSchema(field.Tag.Get("validate"))
				if required && !fuzzy {
					schema.Required = append(schema.Required, configTagName)
				}

				err = mergo.Merge(&fieldSchema, validationSchema, mergo.WithOverride)
				if err != nil {
					return validjsonator.Schema{}, err
				}
			}

			schema.Properties[configTagName] = fieldSchema
		}
	case reflect.Ptr:
		var err error
		schema, err = s.schemaFromType(t.Elem(), fuzzy)
		if err != nil {
			return validjsonator.Schema{}, err
		}
	case reflect.Slice:
		schema.Type = "array"
		itemsSchema, err := s.schemaFromType(t.Elem(), fuzzy)
		if err != nil {
			return validjsonator.Schema{}, err
		}

		schema.Items = &itemsSchema
	case reflect.Array:
		schema.Type = "array"
		itemsSchema, err := s.schemaFromType(t.Elem(), fuzzy)
		if err != nil {
			return validjsonator.Schema{}, err
		}

		schema.Items = &itemsSchema
		schema.MinItems = t.Len()
		schema.MaxItems = t.Len()
	case reflect.Map:
		schema.Type = "object"
		additionalPropertiesSchema, err := s.schemaFromType(t.Elem(), fuzzy)
		if err != nil {
			return validjsonator.Schema{}, err
		}

		schema.AdditionalProperties = &additionalPropertiesSchema
	case reflect.String:
		schema.Type = "string"
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		schema.Type = "integer"
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		schema.Type = "integer"
	case reflect.Float32, reflect.Float64:
		schema.Type = "number"
	case reflect.Bool:
		schema.Type = "boolean"
	case reflect.Interface:
		// Any type, we just leave the schema empty
	default:
		return validjsonator.Schema{}, errors.New("unknown type")
	}

	return schema, nil
}
