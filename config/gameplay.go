package config

type BoardPreset string

const (
	AdvancedBoard     BoardPreset = "ADVANCED"
	IntermediateBoard BoardPreset = "INTERMEDIATE"
	BeginnerBoard     BoardPreset = "BEGINNER"
	CustomBoard       BoardPreset = "CUSTOM"
)

type BoardConfig struct {
	Width      int
	Height     int
	MineCount  int
	LivesCount int
}

const (
	BEGINNER_WIDTH      = 10
	BEGINNER_HEIGHT     = 8
	BEGINNER_MINE_COUNT = 10

	INTERMEDIATE_WIDTH      = 18
	INTERMEDIATE_HEIGHT     = 14
	INTERMEDIATE_MINE_COUNT = 40

	ADVANCED_WIDTH      = 24
	ADVANCED_HEIGHT     = 20
	ADVANCED_MINE_COUNT = 99

	NUM_LIVES = 1
)
