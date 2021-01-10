package service

import (
	"encoding/json"
	"github.com/iurybranco/backend-challenge/discount-calculator/service/database"
	"os"
)

type Config struct {
	LogLevel   int             `json:"logLevel"`
	Database   database.Config `json:"database"`
	ServerPort int             `json:"serverPort"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	var c Config
	return c
}
