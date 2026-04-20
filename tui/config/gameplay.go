package config

type BoardPreset string

const (
	AdvancedBoard     BoardPreset = "ADVANCED"
	IntermediateBoard BoardPreset = "INTERMEDIATE"
	BeginnerBoard     BoardPreset = "BEGINNER"
	CustomBoard       BoardPreset = "CUSTOM"
)

type BoardConfig struct {
	Width     int
	Height    int
	MineCount int
}

const (
	ADVANCED_WIDTH      = 24
	ADVANCED_HEIGHT     = 20
	ADVANCED_MINE_COUNT = 99
	NUM_LIVES           = 1
)
