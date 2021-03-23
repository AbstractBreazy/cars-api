package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	AWAY_STATUS                = "Away"
	IN_STOCK_STATUS            = "In Stock"
	SOLD_OUT_STATUS            = "Sold Out"
	NO_LONGER_AVAILABLE_STATUS = "No Longer Available"
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
