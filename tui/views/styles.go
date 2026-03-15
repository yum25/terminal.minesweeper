package views

import (
	"image/color"

	lipgloss "charm.land/lipgloss/v2"
)

// Colors
var (
	cursorColor = lipgloss.Color("#f8957f")
	mineColor   = lipgloss.Color("#ff7252")

	blue      = lipgloss.Color("#79A1C3")
	green     = lipgloss.Color("#9AC085")
	lightred  = lipgloss.Color("#fc7f90")
	magenta   = lipgloss.Color("#BB8CAF")
	yellow    = lipgloss.Color("#F1CB81")
	cyan      = lipgloss.Color("#77C1D2")
	black     = lipgloss.Color("#394252")
	gray      = lipgloss.Color("#878787")
	lightgray = lipgloss.Color("#cdcdcd")
	white     = lipgloss.Color("#e9e9e9")
)

// Styles
var (
	optionStyle = lipgloss.NewStyle().Padding(0, 2)

	selectedOptionStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Foreground(lipgloss.Color("0")).
				Background(cursorColor).
				Bold(true)

	iconStyle = lipgloss.NewStyle().
			Background(green).
			Padding(0, 1)

	listStyle  = lipgloss.NewStyle().Padding(1)
	titleStyle = lipgloss.NewStyle().Padding(1, 1, 0)

	tileStyle = lipgloss.NewStyle().Width(2).Height(1).AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center)
	boardStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	cursor = tileStyle.Foreground(lipgloss.Color("0")).Background(cursorColor)
)

// Color Maps
var (
	tileColorMap = map[int]color.Color{
		0: black,
		1: blue,
		2: green,
		3: lightred,
		4: magenta,
		5: yellow,
		6: cyan,
		7: black,
		8: gray,
	}

	tileTextColorMap = map[int]color.Color{
		0: black,
		1: black,
		2: black,
		3: black,
		4: black,
		5: black,
		6: black,
		7: white,
		8: black,
	}
)

// Symbols
const (
	flagSymbol      = "⚑"
	falseFlagSymbol = ""
	cursorSymbol    = "<"
)

// Dynamic Styles
func Screen(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
}

func FlaggedStyle(complete bool, mine bool) lipgloss.Style {
	if complete && !mine {
		// Return style for false flagged
		return tileStyle
	}
	return tileStyle.Foreground(black).Background(lightgray)
}

func TileStyle(adjacent int) lipgloss.Style {
	return tileStyle.Foreground(tileTextColorMap[adjacent]).Background(tileColorMap[adjacent])
}
