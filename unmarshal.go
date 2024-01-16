package confuse

import (
	"dario.cat/mergo"
	"os"
)

// Unmarshal unmarshals the configuration files into the given struct.
// It returns an error if the unmarshalling fails.
func (s *Service) Unmarshal(obj any) error {
	fullMap := make(map[string]any)
	for _, source := range s.SourceFiles {
		mappedResult, err := s.unmarshalFile(source)
		if err != nil {
			return err
		}

		err = mergo.Merge(&fullMap, mappedResult, s.MergoConfig...)
		if err != nil {
			return err
		}
	}

	if s.ShouldUseEnvironmentVariables {
		mappedResult, err := s.unmarshalENV(os.Environ())
		if err != nil {
			return err
		}

		err = mergo.Merge(&fullMap, mappedResult, s.MergoConfig...)
	}

	for _, loader := range s.SourceLoaders {
		mappedResult, err := loader()
		if err != nil {
			return err
		}

		err = mergo.Merge(&fullMap, mappedResult, s.MergoConfig...)
		if err != nil {
			return err
		}
	}

	err := parseMapToStruct(obj, fullMap)
	if err != nil {
		return err
	}

	err = s.validate(obj)
	if err != nil {
		return err
	}

	err = s.generateSchemaFiles(obj)
	if err != nil {
		return err
	}

	return nil
}
