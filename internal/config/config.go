package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

const configFileName = ".gatorconfig.json"

func Read() (*Config, error) {
	configPath, err := getConfigFilePath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	cfg := &Config{}
	json.Unmarshal(data, cfg)
	return cfg, nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	write(cfg)
	return nil
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", nil
	}
	return fmt.Sprintf("%s/%s", homePath, configFileName), nil
}

func write(cfg *Config) error {
	configPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	jsonedData, err := json.Marshal(&cfg)
	if err != nil {
		return err
	}
	return os.WriteFile(configPath, jsonedData, 0644)
}
