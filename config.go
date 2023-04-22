package main

import (
	"encoding/json"
	"os"
	"path"
)

type Config struct {
	Token string `json:"token"`
}

func NewConfig() (*Config, error) {
	auth, err := Auth()
	if err != nil {
		return nil, err
	}
	token := auth.Data.SessionToken

	c := Config{Token: token}

	return &c, nil
}

func LoadConfig() (*Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}
	configPath := path.Join(home, ".tastyworks", "config.json")

	_, err = os.Stat(configPath)
	if os.IsNotExist(err) {
		// config file does not exist, create it with NewConfig
		conf, err := NewConfig()
		if err != nil {
			return nil, err
		}

		b, err := json.Marshal(conf)
		if err != nil {
			return nil, err
		}

		err = os.MkdirAll(path.Dir(configPath), 0700)
		if err != nil {
			return nil, err
		}

		err = os.WriteFile(configPath, b, 0600)
		if err != nil {
			return nil, err
		}

		return conf, nil
	} else if err != nil {
		// unexpected error while checking for file existence
		return nil, err
	}

	// config file exists, read and parse it
	b, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
