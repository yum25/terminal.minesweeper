package styles

import (
	"image/color"

	lipgloss "charm.land/lipgloss/v2"
)

// Colors
var (
	CursorColor = lipgloss.Color("#f8957f")
	MineColor   = lipgloss.Color("#fd4343")

	Blue      = lipgloss.Color("#79A1C3")
	Green     = lipgloss.Color("#9AC085")
	Lightred  = lipgloss.Color("#fc7f90")
	Magenta   = lipgloss.Color("#BB8CAF")
	Yellow    = lipgloss.Color("#F1CB81")
	Cyan      = lipgloss.Color("#77c4d2")
	Black     = lipgloss.Color("#394252")
	Gray      = lipgloss.Color("#878787")
	Lightgray = lipgloss.Color("#cdcdcd")
	White     = lipgloss.Color("#e9e9e9")
)

// Color Maps
var (
	tileColorMap = map[int]color.Color{
		0: Lightgray,
		1: Blue,
		2: Green,
		3: Lightred,
		4: Magenta,
		5: Yellow,
		6: Cyan,
		7: Black,
		8: Gray,
	}

	tileTextColorMap = map[int]color.Color{
		0: Lightgray,
		1: Black,
		2: Black,
		3: Black,
		4: Black,
		5: Black,
		6: Black,
		7: White,
		8: Black,
	}
)
