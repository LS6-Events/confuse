package confuse

import "github.com/pelletier/go-toml/v2"

func (s *Service) unmarshalTOML(bytes []byte) (map[string]any, error) {
	var result map[string]any

	err := toml.Unmarshal(bytes, &result)

	return result, err
}
