package main

import "github.com/ls6-events/confuse"

type Config struct {
	Foo            string `validate:"required"`
	Bar            int    `validate:"required"`
	DatabaseConfig DatabaseConfig

	Slice []int
	Map   map[string]int
	Array [3]int
}

type DatabaseConfig struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

func getConfig() (Config, error) {
	var config Config
	err := confuse.New(confuse.WithSourceFiles("./config.yaml"), confuse.WithExactOutputJSONSchema("./schema.json")).Unmarshal(&config)

	return config, err
}
