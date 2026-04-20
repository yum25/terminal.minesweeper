package styles

import (
	"image/color"

	lipgloss "charm.land/lipgloss/v2"
)

// Tailwind-esque defaults
var (
	AlignCenter     = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)
	AlignHorzCenter = lipgloss.NewStyle().AlignHorizontal(lipgloss.Center)
	AlignVertCenter = lipgloss.NewStyle().AlignVertical(lipgloss.Center)
	AlignTop        = lipgloss.NewStyle().AlignVertical(lipgloss.Top)
	AlignBottom     = lipgloss.NewStyle().AlignVertical(lipgloss.Bottom)
	AlignLeft       = lipgloss.NewStyle().AlignHorizontal(lipgloss.Left)
	AlignRight      = lipgloss.NewStyle().AlignHorizontal(lipgloss.Right)

	Padding1  = lipgloss.NewStyle().Padding(1)
	PaddingV1 = lipgloss.NewStyle().Padding(1, 0)
	PaddingH1 = lipgloss.NewStyle().Padding(0, 1)
	PaddingH2 = lipgloss.NewStyle().Padding(0, 2)

	Bold = lipgloss.NewStyle().Bold(true)
)

// Custom Styles
var (
	OptionStyle         = Merge([]lipgloss.Style{AlignCenter, Width(10)})
	DisabledOptionStyle = Merge([]lipgloss.Style{OptionStyle, Text(Gray), Highlight(Darkgray)})
	SelectedOptionStyle = Merge([]lipgloss.Style{OptionStyle, Text(Black), Highlight(CursorColor), Bold})
	IconStyle           = Merge([]lipgloss.Style{Highlight(Green), PaddingH1})
	ListStyle           = Padding1
	TitleStyle          = PaddingV1

	BoardStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	TileStyle  = Merge([]lipgloss.Style{AlignCenter, Width(2), Height(1)})

	FlaggedStyle = Merge([]lipgloss.Style{TileStyle, Text(Lightgray), Highlight(Darkgray)})
	CursorStyle  = Merge([]lipgloss.Style{TileStyle, Text(Black), Highlight(CursorColor)})
	MineStyle    = Merge([]lipgloss.Style{TileStyle, Highlight(Darkgray)})
	MineHitStyle = Merge([]lipgloss.Style{TileStyle, Highlight(Red)})
)

// Helpers

// A tailwind-esque style merge function that inherits styles in a
// cascading manner
func Merge(styles []lipgloss.Style) lipgloss.Style {
	var style lipgloss.Style
	for _, s := range styles {
		style = s.Inherit(style)
	}

	return style
}

func Text(color color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Foreground(color)
}

func Highlight(color color.Color) lipgloss.Style {
	return lipgloss.NewStyle().Background(color)
}

func Width(width int) lipgloss.Style {
	return lipgloss.NewStyle().Width(width)
}

func Height(height int) lipgloss.Style {
	return lipgloss.NewStyle().Height(height)
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
