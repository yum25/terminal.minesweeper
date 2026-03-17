package styles

import lipgloss "charm.land/lipgloss/v2"

// Styles
var (
	OptionStyle         = lipgloss.NewStyle().Width(10).AlignHorizontal(lipgloss.Center)
	SelectedOptionStyle = OptionStyle.Foreground(Black).Background(CursorColor).Bold(true)
	IconStyle           = lipgloss.NewStyle().Background(Green).Padding(0, 1)
	ListStyle           = lipgloss.NewStyle().Padding(1)
	TitleStyle          = lipgloss.NewStyle().Padding(1, 1, 0)

	BoardStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())
	TileStyle  = lipgloss.NewStyle().Width(2).Height(1).AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center)

	FlaggedStyle = TileStyle.Foreground(Lightgray).Background(Darkgray)
	CursorStyle  = TileStyle.Foreground(Black).Background(CursorColor)
	MineStyle    = TileStyle.Background(Darkgray)
	MineHitStyle = TileStyle.Background(MineColor)
)

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
