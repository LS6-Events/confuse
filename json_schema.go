package confuse

import (
	"encoding/json"
	"os"
	"reflect"
)

func (s *Service) generateSchemaFiles(result any) error {
	if s.ExactOutputJSONSchema != "" {
		err := s.generateJSONSchemaFile(result, false, s.ExactOutputJSONSchema)
		if err != nil {
			return err
		}
	}

	if s.FuzzyOutputJSONSchema != "" {
		err := s.generateJSONSchemaFile(result, true, s.FuzzyOutputJSONSchema)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) generateJSONSchemaFile(result any, fuzzy bool, outputFilePath string) error {
	schema, err := s.schemaFromType(reflect.TypeOf(result), fuzzy)
	if err != nil {
		return err
	}

	description := "This is a match of the config struct to the JSON Schema specification."
	if fuzzy {
		description = "This is a fuzzy match of the config struct to the JSON Schema specification."
	}

	jsonSchema := JSONSchema{
		SchemaUri:   "https://json-schema.org/draft/2020-12/schema",
		IDUri:       "http://example.com/schema.json", // TODO: Make this configurable if needed
		Title:       "Confuse Configuration Schema",
		Description: description,
		Schema:      schema,
	}

	bytes, err := json.MarshalIndent(jsonSchema, "", "  ")

	err = os.WriteFile(outputFilePath, bytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
