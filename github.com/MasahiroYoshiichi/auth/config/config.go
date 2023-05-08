package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	AwsRegion  string `json:"aws_region"`
	UserPoolId string `json:"user_pool_id"`
	ClientId   string `json:"client_id"`
	JWtSecret  string `json:"-"`
}

func LoadConfig() (*Config, error) {
	file, err := os.Open("/usr/local/bin/config.json")
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

	config.JWtSecret = os.Getenv("JWT_SECRET_KEY")
	if config.JWtSecret == "" {
		config.JWtSecret = "default_secret_key"
	}

	return config, nil
}
