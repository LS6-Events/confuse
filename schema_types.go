package confuse

type JSONSchema struct {
	SchemaUri   string `json:"$schema"`
	IDUri       string `json:"$id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Schema
}

type Schema struct {
	Ref                  string            `json:"$ref,omitempty"`
	Title                string            `json:"title,omitempty"`
	MultipleOf           float64           `json:"multipleOf,omitempty" `
	Maximum              float64           `json:"maximum,omitempty"`
	ExclusiveMaximum     bool              `json:"exclusiveMaximum,omitempty"`
	Minimum              float64           `json:"minimum,omitempty"`
	ExclusiveMinimum     bool              `json:"exclusiveMinimum,omitempty"`
	MaxLength            int               `json:"maxLength,omitempty"`
	MinLength            int               `json:"minLength,omitempty"`
	Pattern              string            `json:"pattern,omitempty"`
	MaxItems             int               `json:"maxItems,omitempty"`
	MinItems             int               `json:"minItems,omitempty"`
	UniqueItems          bool              `json:"uniqueItems,omitempty"`
	MaxProperties        int               `json:"maxProperties,omitempty"`
	MinProperties        int               `json:"minProperties,omitempty"`
	Required             []string          `json:"required,omitempty"`
	Enum                 []interface{}     `json:"enum,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Format               string            `json:"format,omitempty"`
	AllOf                []Schema          `json:"allOf,omitempty"`
	OneOf                []Schema          `json:"oneOf,omitempty"`
	AnyOf                []Schema          `json:"anyOf,omitempty"`
	Not                  *Schema           `json:"not,omitempty"`
	Items                *Schema           `json:"items,omitempty"`
	Properties           map[string]Schema `json:"properties,omitempty"`
	PatternProperties    map[string]Schema `json:"patternProperties,omitempty"`
	AdditionalProperties *Schema           `json:"additionalProperties,omitempty"`
	Description          string            `json:"description,omitempty"`
}
