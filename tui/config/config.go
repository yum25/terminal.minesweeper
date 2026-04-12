package config

import (
	"encoding/json"
	"os"
	"path/filepath"

	"terminal.minesweeper/tui/constants"
)

func appDir() (string, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}
	dir := filepath.Join(configDir, "terminal.minesweeper")

	err = os.MkdirAll(dir, 0755)
	return dir, err
}

func statsPath() (string, error) {
	dir, err := appDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "stats.json"), nil
}

func configPath() (string, error) {
	dir, err := appDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(dir, "config.json"), nil
}

func LoadConfig() (*constants.Config, error) {
	path, err := configPath()
	if err != nil {
		return nil, err
	}
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &constants.DEFAULT_CONFIG, nil // default config
	}

	var config constants.Config
	err = json.Unmarshal(data, &config)

	return &config, err
}

func SaveConfig(config constants.Config) error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadStats() (constants.Stats, error) {
	path, err := statsPath()
	if err != nil {
		return constants.Stats{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return constants.Stats{}, nil // first run, no stats yet
	}

	var stats constants.Stats
	err = json.Unmarshal(data, &stats)
	return stats, err
}

func SaveStats(stats constants.Stats) error {
	path, err := statsPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(stats)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
