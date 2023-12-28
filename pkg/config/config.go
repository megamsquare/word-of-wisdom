package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type ConfigKey string

const ConfigCtxKey ConfigKey = "config"

type Config struct {
	ServerHost            string `envconfig:"SERVER_HOST"`
	ServerPort            int    `envconfig:"SERVER_PORT"`
	HashcashZerosCount    int
	HashcashDuration      int64
	HashcashMaxLoop int
}

func Load(path string) (*Config, error) {
	config := Config{}
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	decode := json.NewDecoder(file)
	err = decode.Decode(&config)
	if err != nil {
		return &config, err
	}
	err = envconfig.Process("", &config)
	return &config, err
}
