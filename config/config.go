package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)


// DatabaseConfigurations exported
type DBConfig struct {
	Engine   string
	Server   string
	Port     string
	User     string
	Password string
	Database string
}

type Config struct {
	DatabaseConfiguration DBConfig
}

// Reading DataConfigurations file
func GetConfig() (Config, error) {
	config := Config{}
	file, err := os.Open("./config/config.json")
	if err != nil {
		return config, err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return config, err
	}
	return config, nil
}
