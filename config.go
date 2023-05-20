package tasty

import (
	"encoding/json"
	"os"
	"path"
	"time"
)

type Config struct {
	Token     string    `json:"token"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewConfig() (*Config, error) {
	auth, err := Auth()
	if err != nil {
		return nil, err
	}

	c := Config{
		Token:     auth.Data.SessionToken,
		UpdatedAt: time.Now().UTC(),
	}

	return &c, nil
}

func writeConfig(configPath string) (*Config, error) {
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
		return writeConfig(configPath)
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

	expiresAfter := 12 * time.Hour
	if time.Now().UTC().After(c.UpdatedAt.Add(expiresAfter)) {
		return writeConfig(configPath)
	}

	return &c, nil
}
