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

func DefaultConfigReader() ConfigReader {
	return ConfigReader{
		File:        "./properties.json",
		Environment: true,
	}
}

func (c ConfigReader) Read(props interface{}) error {
	log.Printf("Reading properties from %s", c.File)
	configFile, err := ioutil.ReadFile(c.File)
	if err != nil {
		log.Printf("Error reading configuration file %s. %v", c.File, err)
		return err
	}

	err = json.Unmarshal(configFile, props)
	if err != nil {
		log.Printf("Error unmarshalling configuration file into given interface %v", props)
		return err
	}

	// TODO read from environment and override values from config

	return nil
}
