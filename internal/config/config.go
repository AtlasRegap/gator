package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// The file name only. The home folder gets added in front of it below.
const configFileName = ".gatorconfig.json"

// One field per JSON key. The struct tag maps each key to its field.
// The fields are capitalised so the json package can see them.
type Config struct {
	DBURL           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

// Builds the full path, e.g. /home/sagar/.gatorconfig.json
func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, configFileName), nil
}

// Reads the file: path → raw bytes → Config struct
func Read() (Config, error) {
	path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil { // &cfg so Unmarshal fills in YOUR variable
		return Config{}, err
	}
	return cfg, nil
}

// Writes the file: Config struct → JSON bytes → file (replaces what's there)
func write(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(cfg, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, data, 0600) // 0600 = only you can read/write it
}

// A method on Config. The pointer receiver (*Config) means it changes
// the caller's real struct, not a copy. Then it saves to disk.
func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUserName = username
	return write(*cfg)
}
