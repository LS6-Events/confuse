package confuse

import (
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestService_Validate(t *testing.T) {
	type testStruct struct {
		Hello string `validate:"required"`
	}

	t.Run("should not validate if ShouldValidate is false", func(t *testing.T) {
		s := Service{
			ShouldValidate: false,
		}
		err := s.validate(testStruct{})
		require.NoError(t, err)
	})

	t.Run("should not validate if Validator is nil", func(t *testing.T) {
		s := Service{
			ShouldValidate: true,
			Validator:      nil,
		}
		err := s.validate(testStruct{})
		require.NoError(t, err)
	})

	t.Run("should validate if ShouldValidate is true and Validator is not nil", func(t *testing.T) {
		s := Service{
			ShouldValidate: true,
			Validator:      validator.New(),
		}
		err := s.validate(testStruct{})
		require.Error(t, err)
	})
}
