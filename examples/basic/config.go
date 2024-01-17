package main

import (
	"github.com/ls6-events/confuse"
	"time"
)

type Config struct {
	Foo            string `validate:"required"`
	Bar            int    `validate:"required"`
	DatabaseConfig DatabaseConfig

	Slice []int
	Map   map[string]int
	Array [3]int

	DiffName DiffName `config:"different_name"`

	PointerStruct *PointerStruct

	Date time.Time
}

type DatabaseConfig struct {
	Host string `validate:"required"`
	Port int    `validate:"required"`
}

type DiffName struct {
	Hello string
}

type PointerStruct struct {
	Hello string
}

func getConfig() (Config, error) {
	var config Config
	err := confuse.New(confuse.WithSourceFiles("./config.yaml")).Unmarshal(&config)

	return config, err
}
