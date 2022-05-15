package config_reader

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type ConfigReader struct {
	File        string
	Environment bool
}

func NewConfigReader() ConfigReader {
	return ConfigReader{
		File:        "./properties.json",
		Environment: true,
	}
}

func (c ConfigReader) Read(props interface{}) {
	configFile, err := ioutil.ReadFile(c.File)
	if err != nil {
		log.Fatalf("Error reading configuration file %s. %v", c.File, err)
	}
	err = json.Unmarshal(configFile, props)
	if err != nil {
		panic(err)
	}
}
