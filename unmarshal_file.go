package confuse

import (
	"os"
	"path/filepath"
)

func (s *Service) unmarshalFile(filePath string) (map[string]any, error) {
	if _, err := os.Stat(filePath); err != nil {
		return nil, err
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var result map[string]any
	switch filepath.Ext(filePath) {
	case ".json":
		result, err = s.unmarshalJSON(bytes)
	case ".yaml", ".yml":
		result, err = s.unmarshalYAML(bytes)
	case ".toml":
		result, err = s.unmarshalTOML(bytes)
	case ".env":
		result, err = s.unmarshalENVFile(bytes)
	}

	return result, err
}
