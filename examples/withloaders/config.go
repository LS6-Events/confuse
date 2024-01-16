package main

import "github.com/ls6-events/confuse"

type Config struct {
	DatabaseConfig DatabaseConfig
}

type DatabaseConfig struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

func getConfig() (Config, error) {
	var config Config
	err := confuse.New(confuse.WithSourceFiles("./config.yaml"), confuse.WithSourceLoaders(func() (map[string]any, error) {
		return map[string]any{
			"database_config": map[string]any{
				"host": "localhost",
			},
		}, nil
	})).Unmarshal(&config)

	return config, err
}
