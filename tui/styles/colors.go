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
	Darkgray  = lipgloss.Color("#394252")
	Gray      = lipgloss.Color("#878787")
	Lightgray = lipgloss.Color("#cdcdcd")
	White     = lipgloss.Color("#e9e9e9")
	Black     = lipgloss.Color("0")
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
		7: Darkgray,
		8: Gray,
	}

	tileTextColorMap = map[int]color.Color{
		0: Lightgray,
		1: Darkgray,
		2: Darkgray,
		3: Darkgray,
		4: Darkgray,
		5: Darkgray,
		6: Darkgray,
		7: White,
		8: Darkgray,
	}
)
