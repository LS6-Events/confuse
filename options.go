package confuse

import (
	"dario.cat/mergo"
	"github.com/go-playground/validator/v10"
)

type Option func(*Service)

// WithSourceFiles sets the source files to read from.
// It can be used multiple times to read from multiple files.
// The first file in the list takes the lowest precedence, and the last file takes the highest precedence.
// E.g. if the same key is defined in both files, the value in the last file will be used.
// Base -> Override 1 -> Override 2 -> Override 3
func WithSourceFiles(files ...string) Option {
	return func(s *Service) {
		s.SourceFiles = append(s.SourceFiles, files...)
	}
}

// WithExactOutputJSONSchema sets the path to the output JSON schema file that is generated from the unmarshalled configuration.
// If this is set, the output JSON schema will be written to this file.
func WithExactOutputJSONSchema(path string) Option {
	return func(s *Service) {
		s.ExactOutputJSONSchema = path
	}
}

// WithFuzzyOutputJSONSchema sets the path to the output JSON schema file that is generated from the unmarshalled configuration.
// It is similar to ExactOutputJSONSchema, but all the required fields are made optional.
// This is useful when you want to use the override files mechanism to override only a subset of the configuration.
// If this is set, the output JSON schema will be written to this file.
func WithFuzzyOutputJSONSchema(path string) Option {
	return func(s *Service) {
		s.FuzzyOutputJSONSchema = path
	}
}

// WithJSONSchemaKeyModifier sets the function to use to modify the JSON schema key.
// If this is set, the function will be used to modify the JSON schema key.
// By default, it is set to strcase.ToSnake.
func WithJSONSchemaKeyModifier(modifier func(string) string) Option {
	return func(s *Service) {
		s.JSONSchemaKeyModifier = modifier
	}
}

// WithValidation sets the flag to indicate whether the unmarshalled configuration should be validated against the JSON schema.
// If this is set to true, the unmarshalled configuration will be validated against the JSON schema.
func WithValidation(shouldValidate bool) Option {
	return func(s *Service) {
		s.ShouldValidate = shouldValidate
	}
}

// WithCustomValidator sets the validator to use to validate the unmarshalled configuration.
// If this is not set, the default will be used will be used.
func WithCustomValidator(validator *validator.Validate) Option {
	return func(s *Service) {
		s.ShouldValidate = true
		s.Validator = validator
	}
}

// WithEnvironmentVariables sets the flag to indicate whether to use the environment variables to override the configuration.
// If this is set to true, the environment variables will be used to override the configuration.
func WithEnvironmentVariables(shouldUseEnvironmentVariables bool) Option {
	return func(s *Service) {
		s.ShouldUseEnvironmentVariables = shouldUseEnvironmentVariables
	}
}

// WithEnvironmentVariablesPrefix sets the prefix of the environment variables to use to override the configuration.
// If this is set, only the environment variables with the prefix will be used to override the configuration.
func WithEnvironmentVariablesPrefix(prefix string) Option {
	return func(s *Service) {
		s.EnvironmentVariablesPrefix = prefix
	}
}

// WithEnvironmentVariablesSeparator sets the separator of the environment variables to use to override the configuration.
// By default, the separator is "__".
func WithEnvironmentVariablesSeparator(separator string) Option {
	return func(s *Service) {
		s.EnvironmentVariablesSeparator = separator
	}
}

// WithMergoConfig sets the list of options to use when merging the configuration using dario.cat/mergo.
// By default it just uses mergo.WithOverride.
func WithMergoConfig(config ...func(*mergo.Config)) Option {
	return func(s *Service) {
		s.MergoConfig = config
	}
}
