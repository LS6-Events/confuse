package confuse

import (
	"dario.cat/mergo"
	"github.com/go-playground/validator/v10"
	"github.com/iancoleman/strcase"
)

func New(opts ...Option) *Service {
	s := &Service{
		EnvironmentVariablesSeparator: "__",
		JSONSchemaKeyModifier:         strcase.ToSnake,
	}

	for _, opt := range opts {
		opt(s)
	}

	if s.ShouldValidate && s.Validator == nil {
		s.Validator = validator.New(validator.WithRequiredStructEnabled())
	}
	if len(s.MergoConfig) == 0 {
		s.MergoConfig = []func(*mergo.Config){mergo.WithOverride}
	}

	return s
}
