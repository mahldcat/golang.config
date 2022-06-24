package config

import (
	"errors"
	"fmt"
)

func GetConfigEntry(key string) (string, error) {
	if key == "" {
		return "", errors.New("empty string")
	}

	config := fmt.Sprintf("Fetch config entry value for Key:= %v", key)

	return config, nil

}
