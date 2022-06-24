package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Configuration struct {
	Connection string
}

//GetConfiguration will read in configuraiton data based off a file "configuration.json"
func GetConfiguration() (Configuration, error) {
	return getConfigImpl("configuration.json")
}

//GetConfiguration will read in configuration data based off the supplied path argument
func GetConfigurationNamedFile(path string) (Configuration, error) {
	return getConfigImpl(path)
}

func getConfigImpl(path string) (Configuration, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		return Configuration{}, err
	}

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)

	return configuration, err
}

//GetConfigEntry will return an entry
//TODO: Research generic structs?
func GetConfigEntry(conf *Configuration, key string) (string, error) {
	if key == "" {
		return "", errors.New("empty string")
	}

	c := fmt.Sprintf("Key[%v]:=", key)

	return c, nil
}
