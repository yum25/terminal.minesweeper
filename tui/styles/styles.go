package styles

import (
	"image/color"
	"slices"

	lipgloss "charm.land/lipgloss/v2"
)

// Tailwind-esque defaults
var (
	AlignCenter     = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	AlignHorzCenter = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
	AlignVertCenter = lipgloss.NewStyle().AlignVertical(lipgloss.Center)
	AlignHorzLeft   = lipgloss.NewStyle().AlignHorizontal(lipgloss.Left)
	AlignHorzRight  = lipgloss.NewStyle().AlignHorizontal(lipgloss.Right)

	PaddingH1 = lipgloss.NewStyle().Padding(0, 1)
	PaddingH2 = lipgloss.NewStyle().Padding(0, 2)
)

// Custom Styles
var (
	OptionStyle         = lipgloss.NewStyle().Width(10).AlignHorizontal(lipgloss.Center)
	DisabledOptionStyle = OptionStyle.Foreground(Gray).Background(Darkgray)
	SelectedOptionStyle = OptionStyle.Foreground(Black).Background(CursorColor).Bold(true)
	IconStyle           = lipgloss.NewStyle().Background(Green).Padding(0, 1)
	ListStyle           = lipgloss.NewStyle().Padding(1)
	TitleStyle          = lipgloss.NewStyle().Padding(1, 1, 0)

	BoardStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	TileStyle  = AlignCenter.Width(2).Height(1)

	FlaggedStyle = TileStyle.Foreground(Lightgray).Background(Darkgray)
	CursorStyle  = TileStyle.Foreground(Black).Background(CursorColor)
	MineStyle    = TileStyle.Background(Darkgray)
	MineHitStyle = TileStyle.Background(Red)
)

// Helpers

// A tailwind-esque style merge function that inherits styles starting from the
// end of the array, imitating Cascading Style Sheets
func Merge(styles []lipgloss.Style) lipgloss.Style {
	var style lipgloss.Style
	for _, s := range slices.Backward(styles) {
		style.Inherit(s)
	}

	return style
}

func Text(color color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color)
}

func Highlight(color color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Background(color)
}

// Dynamic Styles
func Screen(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
}

func RevealedStyle(adjacent int) lipgloss.Style {
	return TileStyle.Foreground(tileTextColorMap[adjacent]).Background(tileColorMap[adjacent])
}
