# Confuse
Confuse is a configuration library for Golang that is built on the idea of using a struct to manage your configuration data. This struct can then be populated with a _fusion_ of configuration from any number of configuration sources such as environment variables, configuration files, and more.

The name Confuse was a merging between _config_ and _fusion_, to indicate that two or more configuration files can 
be "fused" together in an override specification.

It is designed to be a lightweight, strongly typed and easy to use alternative to using `map[string]interface{}` to load configuration data, or accessing fields using `.Get()` from [viper](https://www.github.com/spf13/viper) and parsing the types directly.

## Features
- Load configuration from `yaml`, `json`, `toml` and `env` files.
- Load override configuration from environment variables.
- Load configuration from multiple sources, that will override each other in order of precedence.
- Load configuration into a struct, with specific data types for each field.
- Allow usage for `go-validator` to validate configuration data (`validate:"required"`).
- Generate a `json-schema` from your configuration struct to use with `.json` and `.yaml` files.

## Installation
```bash
go get github.com/ls6-events/confuse
```

## Usage
### Basic

```go
package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/ls6-events/confuse"
)

type Config struct {
	// `config` tag is not required, we will check for the field name in it's case, snake_case, kebab-case, camelCase and PascalCase forms.
	Name string `config:"name"`
	Port int    `config:"port"`

	// You can also use the `validate` tag to validate the configuration data.
	Host string `config:"host" validate:"required"`

	// You can also use nested structs.
	Database struct {
		Host string `config:"host"`
		Port int    `config:"port"`
	} `config:"database"`
}

func main() {
	// Create a new configuration struct.
	var config Config

	// Load the configuration from the default sources.
	err := confuse.New(confuse.WithSourceFiles("./config.yaml")).Unmarshal(&config)

	// Check for errors.
	if err != nil {
		panic(err)
	}

	// Print the configuration. 
	fmt.Printf("Name: %s\n", config.Name)
	fmt.Printf("Port: %d\n", config.Port)
	fmt.Printf("Host: %s\n", config.Host)

	// Print the nested configuration.
	spew.Dump(config.Database)
}
```

Check out more from the [examples](./examples) directory for different use cases.

### Multiple Sources
There is an example of using multiple sources in the [examples](./examples/withoverrides) directory, but here is a quick example.

```go
package main

import (
    "fmt"
    "github.com/ls6-events/confuse"
)

type Config struct {
    Name string `config:"name"`
    Port int    `config:"port"`
}

func main() {
    // Create a new configuration struct.
    var config Config

    // Load the configuration from the default sources.
    err := confuse.New(
        confuse.WithSourceFiles("./config.yaml", "./override.json"),
    ).Unmarshal(&config)

    // Check for errors.
    if err != nil {
        panic(err)
    }

    // Print the configuration. 
    fmt.Printf("Name: %s\n", config.Name) // override
    fmt.Printf("Port: %d\n", config.Port) // 8080
}
```
```yaml
# config.yaml
name: "config"
port: 8080
```

```json
// override.json
{
  // This will override the name from the config.yaml file.
    "name": "override"
}
```

You can have as many sources as you want, and they will be loaded in order of precedence. The last source will override the previous sources.

### Environment Variables
You can also load configuration from environment variables. This is done by using the `confuse.WithSourceEnv()` option.

```go
package main

import (
    "fmt"
    "github.com/ls6-events/confuse"
)

type Config struct {
    Name string `config:"name"`
    Port int    `config:"port"`
}

func main() {
    // Create a new configuration struct.
    var config Config

    // Load the configuration from the default sources.
    err := confuse.New(
		confuse.WithSourceFiles("./config.yaml"), 
		confuse.WithEnvironmentVariables(true),
    ).Unmarshal(&config)

    // Check for errors.
    if err != nil {
        panic(err)
    }

    // Print the configuration. 
    fmt.Printf("Name: %s\n", config.Name) // override
    fmt.Printf("Port: %d\n", config.Port) // 8080
}
```

This will allow the env variable `NAME` to override the `name` field in the configuration struct. The env variable `PORT` will override the `port` field in the configuration struct. If there is a nested struct, the joining operator of `__` will be used to join the field names together. If you want to change this, there is the option `confuse.WithEnvironmentVariablesSeparator(separator string)` function to change it.  

### Validation
You can also use the `validate` tag to validate the configuration data. This is done by using the `confuse.WithValidation(true)` option.

```go
package main

import (
    "fmt"
    "github.com/ls6-events/confuse"
)

type Config struct {
    Name string `config:"name" validate:"required"`
    Port int    `config:"port" validate:"required"`
}

func main() {
    // Create a new configuration struct.
    var config Config

    // Load the configuration from the default sources.
    err := confuse.New(
        confuse.WithSourceFiles("./config.yaml"), 
        confuse.WithValidation(true),
    ).Unmarshal(&config)

    // Check for errors.
    if err != nil {
        panic(err)
    }

    // Print the configuration. 
    fmt.Printf("Name: %s\n", config.Name) // override
    fmt.Printf("Port: %d\n", config.Port) // 8080
}
```

There is an example of using validation in the [examples](./examples/withvalidation) directory. It will also use these tags to generate the correct JSON schema where appropriate.


### JSON Schema

You can also generate a JSON schema from your configuration struct. This is done by using the `confuse.WithExactOutputJSONSchema(filepath string)` option.

```go
package main

import (
    "fmt"
    "github.com/ls6-events/confuse"
)

type Config struct {
    Name string `config:"name" validate:"required"`
    Port int    `config:"port" validate:"required"`
}

func main() {
    // Create a new configuration struct.
    var config Config

    // Load the configuration from the default sources.
    err := confuse.New(
        confuse.WithSourceFiles("./config.yaml"), 
        confuse.WithExactOutputJSONSchema("./schema.json"),
    ).Unmarshal(&config)

    // Check for errors.
    if err != nil {
        panic(err)
    }

    // Print the configuration. 
    fmt.Printf("Name: %s\n", config.Name) // override
    fmt.Printf("Port: %d\n", config.Port) // 8080
}
```

There is an example of using validation in the [examples](./examples/withjsonschema) directory. It will also use these tags to generate the correct JSON schema where appropriate.

Note, we also have an option to create a _fuzzy_ schema, which will not include the `required` tags (it's useful for setting the schema of override files as the nature of them may indicate that not all properties are set). This is done by using the `confuse.WithFuzzyOutputJSONSchema(filepath string)` option.

### Custom Sources

You can also create your own custom sources. This can be done by using the `confuse.WithSourceLoaders(func() (map[string]any, error))` option.

**Note: this step happens before the unmarshalling to a struct. Therefore if you want it to override any existing configuration keys, you must preserve the case of the keys.**

```go
package main

import (
    "fmt"
    "github.com/ls6-events/confuse"
)

type Config struct {
    Name string `config:"name"`
    Port int    `config:"port"`
}

func main() {
    // Create a new configuration struct.
    var config Config

    // Load the configuration from the default sources.
    err := confuse.New(
        confuse.WithSourceFiles("./config.yaml"), 
        confuse.WithSourceLoaders(func() (map[string]any, error) {
            return map[string]any{
                "name": "override",
            }, nil
        }),
    ).Unmarshal(&config)

    // Check for errors.
    if err != nil {
        panic(err)
    }

    // Print the configuration. 
    fmt.Printf("Name: %s\n", config.Name) // override
    fmt.Printf("Port: %d\n", config.Port) // 8080
}
```

There is an example of using custom sources in the [examples](./examples/withloaders) directory.
