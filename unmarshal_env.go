package confuse

import (
	"github.com/iancoleman/strcase"
	"strings"
)

func (s *Service) unmarshalENVFile(bytes []byte) (map[string]any, error) {
	envVars := strings.Split(string(bytes), "\n")

	return s.unmarshalENV(envVars)
}

func (s *Service) unmarshalENV(envVars []string) (map[string]any, error) {
	result := make(map[string]any)
	for _, envVar := range envVars {
		split := strings.Split(envVar, "=")
		key := split[0]
		value := split[1]

		if s.EnvironmentVariablesPrefix != "" && !strings.HasPrefix(key, s.EnvironmentVariablesPrefix) {
			continue
		}

		key = strings.TrimPrefix(key, s.EnvironmentVariablesPrefix)
		key = strcase.ToSnake(key)
		keyPath := strings.Split(key, s.EnvironmentVariablesSeparator)

		var currentMap = result
		for i, key := range keyPath {
			if i == len(keyPath)-1 {
				currentMap[key] = value
				continue
			}

			if _, ok := currentMap[key]; !ok {
				currentMap[key] = make(map[string]any)
			}

			currentMap = currentMap[key].(map[string]any)
		}
	}

	return result, nil
}
