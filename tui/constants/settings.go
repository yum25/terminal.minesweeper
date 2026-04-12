package constants

type Config struct {
	Width     int `json:"width"`
	Height    int `json:"height"`
	MineCount int `json:"mine_count"`
}

type Stats struct {
	GamesPlayed int `json:"games_played"`
	GamesWon    int `json:"games_won"`
}

const (
	ADVANCED_WIDTH      = 24
	ADVANCED_HEIGHT     = 20
	ADVANCED_MINE_COUNT = 99
	NUM_LIVES           = 1
)

var (
	Width     = ADVANCED_WIDTH
	Height    = ADVANCED_HEIGHT
	MineCount = ADVANCED_MINE_COUNT
	Lives     = NUM_LIVES
)

var DEFAULT_CONFIG = Config{
	Width:     ADVANCED_WIDTH,
	Height:    ADVANCED_HEIGHT,
	MineCount: ADVANCED_MINE_COUNT,
}
