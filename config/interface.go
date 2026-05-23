package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	Board        BoardConfig `json:"board"`
	BoardType    BoardPreset `json:"board_type"`
	UserControls UserControlsMap
	GameControls GameControlsMap
	ControlType  ControlPreset `json:"control_type"`
}

type Stats struct {
	GamesPlayed int `json:"games_played"`
	GamesWon    int `json:"games_won"`
}

var DEFAULT_USERKEYMAP, DEFAULT_GAMEKEYMAP = DEFAULT_CONTROLS.ToKeyMap()
var DEFAULT_CONFIG = Config{
	Board: BoardConfig{
		Width:      ADVANCED_WIDTH,
		Height:     ADVANCED_HEIGHT,
		MineCount:  ADVANCED_MINE_COUNT,
		LivesCount: 1,
	},
	BoardType:    AdvancedBoard,
	UserControls: DEFAULT_USERKEYMAP,
	GameControls: DEFAULT_GAMEKEYMAP,
	ControlType:  DefaultControls,
}

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

func (c *Config) LoadConfig() error {
	path, err := configPath()
	if err != nil {
		return err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var local Config
	err = json.Unmarshal(data, &local)

	*c = local
	return err
}

func (c *Config) SaveConfig() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}

func LoadStats() (*Stats, error) {
	path, err := statsPath()
	if err != nil {
		return &Stats{}, err
	}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return &Stats{}, nil // first run, no stats yet
	}

	var stats Stats
	err = json.Unmarshal(data, &stats)
	return &stats, err
}

func SaveStats(stats *Stats) error {
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

// Hooks into standard json.Marshal and json.Unmarshal functions
func (c *Config) UnmarshalJSON(data []byte) error {
	type ConfigAlias Config
	aux := &struct {
		RawControls Controls `json:"controls"`
		*ConfigAlias
	}{
		ConfigAlias: (*ConfigAlias)(c),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	c.UserControls, c.GameControls = aux.RawControls.ToKeyMap()
	return nil
}

func (c *Config) MarshalJSON() ([]byte, error) {
	type ConfigAlias Config
	return json.Marshal(&struct {
		RawControls Controls `json:"controls"`
		*ConfigAlias
	}{
		RawControls: c.FromKeyMap(),
		ConfigAlias: (*ConfigAlias)(c),
	})
}
