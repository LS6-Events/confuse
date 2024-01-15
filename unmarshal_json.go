package confuse

import (
	jsoniter "github.com/json-iterator/go"
)

func (s *Service) unmarshalJSON(bytes []byte) (map[string]any, error) {
	var result map[string]any

	err := jsoniter.Unmarshal(bytes, &result)

	return result, err
}
