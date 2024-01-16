package confuse

import (
	"dario.cat/mergo"
	"github.com/go-playground/validator/v10"
)

// Service is the configuration for confuse, which is used to unmarshal config files
type Service struct {
	// SourceFiles is the list of files to load.
	// The first file in the list takes the lowest precedence, and the last file takes the highest precedence.
	// E.g. if the same key is defined in both files, the value in the last file will be used.
	// Base -> Override 1 -> Override 2 -> Override 3
	SourceFiles []string

	// SourceLoaders is the list of loaders to use to load the files.
	// The first loader in the list takes the lowest precedence, and the last loader takes the highest precedence.
	// They take a higher precedence than SourceFiles.
	// E.g. if the same key is defined in both files, the value in the last loader will be used.
	// Base -> Override 1 -> Override 2 -> Override 3
	SourceLoaders []Loader

	// ExactOutputJSONSchema is the path to the output JSON schema file that is generated from the unmarshalled configuration.
	// If this is set, the output JSON schema will be written to this file.
	ExactOutputJSONSchema string

	// FuzzyOutputJSONSchema is the path to the output JSON schema file that is generated from the unmarshalled configuration.
	// It is similar to ExactOutputJSONSchema, but all the required fields are made optional.
	// This is useful when you want to use the override files mechanism to override only a subset of the configuration.
	// If this is set, the output JSON schema will be written to this file.
	FuzzyOutputJSONSchema string

	// JSONSchemaKeyModifier is the function to use to modify the JSON schema key.
	// If this is set, the function will be used to modify the JSON schema key.
	// By default, it is set to strcase.ToSnake.
	JSONSchemaKeyModifier func(string) string

	// ShouldValidate is a flag to indicate whether the unmarshalled configuration should be validated against the JSON schema.
	// If this is set to true, the unmarshalled configuration will be validated against the JSON schema.
	ShouldValidate bool

	// Validator is the validator to use to validate the unmarshalled configuration.
	// If this is not set, nil will be used.
	Validator *validator.Validate

	// ShouldUseEnvironmentVariables is a flag to indicate whether to use the environment variables to override the configuration.
	// If this is set to true, the environment variables will be used to override the configuration.
	ShouldUseEnvironmentVariables bool

	// EnvironmentVariablesPrefix is the prefix of the environment variables to use to override the configuration.
	// If this is set, only the environment variables with the prefix will be used to override the configuration.
	EnvironmentVariablesPrefix string

	// EnvironmentVariablesSeparator is the separator of the environment variables to use to override the configuration.
	// By default, the separator is "__".
	EnvironmentVariablesSeparator string

	// MergoConfig is the list of options to use when merging the configuration using dario.cat/mergo.
	// By default it just uses mergo.WithOverride.
	MergoConfig []func(*mergo.Config)
}
