package main

import "github.com/ls6-events/confuse"

type Config struct {
	Foo string `validate:"required,oneof=foo bar"`
	Bar int    `validate:"required,min=1,max=10"`
}

func getConfig(fail bool) (Config, error) {
	var config Config

	sourceFile := "./config.yaml"
	if fail {
		// To show the error message, we need to use a different source file that will fail.
		sourceFile = "./fail_config.yaml"
	}

	err := confuse.New(confuse.WithSourceFiles(sourceFile), confuse.WithExactOutputJSONSchema("./schema.json"), confuse.WithValidation(true)).Unmarshal(&config)

	return config, err
}
