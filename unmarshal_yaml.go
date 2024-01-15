package confuse

import (
	"gopkg.in/yaml.v3"
)

func (s *Service) unmarshalYAML(bytes []byte) (map[string]any, error) {
	var result map[string]any

	err := yaml.Unmarshal(bytes, &result)

	return result, err
}
