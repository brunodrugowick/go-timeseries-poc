package config

import "fmt"

type Config struct {
	locations   []string
	environment bool
}

func NewConfig() Config {
	return Config{
		// default location is current directory
		locations: []string{"."},
		// override with corresponding values from environemnt, if existent
		environment: true,
	}
}

func (c Config) ReadEnvironment(r bool) {
	c.environment = r
}

func (c Config) AddLocation(l string) {
	c.locations = append(c.locations, l)
}

func (c Config) Read(properties *interface{}) map[string][]string {
	for _, location := range c.locations {
		fmt.Printf("Reading config from %s", location)
	}

	return map[string][]string{}
}
