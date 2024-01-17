package main

import (
	"github.com/davecgh/go-spew/spew"
)

func main() {
	config, err := getConfig()
	if err != nil {
		panic(err)
	}

	spew.Dump(config)
}
