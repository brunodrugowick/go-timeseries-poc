package config_reader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
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

	// TODO organize this test code
	if c.Environment {
		v := reflect.TypeOf(props)
		if v.Kind() == reflect.Pointer {
			v = v.Elem()
		}
		var tags []string
		getTags(&tags, v, "")
		fmt.Println(tags)
	}

	return nil
}

func getTags(tags *[]string, t reflect.Type, prefix string) {
	if (t.Kind() > reflect.Bool && t.Kind() <= reflect.Array) || (t.Kind() == reflect.String) {
		*tags = append(*tags, prefix)
		prefix = strings.ToUpper(prefix)
		if env, ok := os.LookupEnv(prefix); ok {
			// TODO I know what to override and with what... but don't now how to do it.
			log.Printf("Need to override value %s with %s", prefix, env)
		}
		return
	}

	for i := 0; i < t.NumField(); i++ {
		var tag string
		tag = t.Field(i).Tag.Get("json")

		if prefix == "" {
			getTags(tags, t.Field(i).Type, tag)
		} else {
			getTags(tags, t.Field(i).Type, prefix+"_"+tag)
		}
	}
}
