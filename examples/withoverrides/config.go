package main

import "github.com/ls6-events/confuse"

type Config struct {
	Foo            string `validate:"required"`
	Bar            int    `validate:"required"`
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

func getConfig() (Config, error) {
	var config Config
	err := confuse.New(confuse.WithSourceFiles("./config.yaml", "./config2.yaml", "./config3.yaml")).Unmarshal(&config)

	return config, err
}
