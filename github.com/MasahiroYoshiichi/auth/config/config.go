package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AwsRegion  string `json:"aws_region"`
	UserPoolId string `json:"user_pool_id"`
	ClientId   string `json:"client_id"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("config/config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := &Config{}
	err = decoder.Decode(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
