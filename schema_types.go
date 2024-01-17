package confuse

import "github.com/ls6-events/validjsonator"

type JSONSchema struct {
	SchemaUri   string `json:"$schema"`
	IDUri       string `json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	validjsonator.Schema
}
