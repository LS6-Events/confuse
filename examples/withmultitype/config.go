package main

import "github.com/ls6-events/confuse"

type Config struct {
	Json JsonConfig
	Yaml YamlConfig
	Toml TomlConfig
	Env  EnvConfig
}

type JsonConfig struct {
	Json bool
}

type YamlConfig struct {
	Yaml bool
}

type TomlConfig struct {
	Toml bool
}

type EnvConfig struct {
	Env bool
}

func getConfig() (Config, error) {
	var config Config
	err := confuse.New(confuse.WithSourceFiles("./config.env", "./config.json", "./config.toml", "./config.yaml")).Unmarshal(&config)

	return config, err
}
