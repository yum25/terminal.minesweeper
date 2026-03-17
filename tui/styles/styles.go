package styles

import lipgloss "charm.land/lipgloss/v2"

// Styles
var (
	OptionStyle = lipgloss.NewStyle().Padding(0, 2)

	SelectedOptionStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Foreground(lipgloss.Color("0")).
				Background(CursorColor).
				Bold(true)

	IconStyle = lipgloss.NewStyle().
			Background(Green).
			Padding(0, 1)

	ListStyle  = lipgloss.NewStyle().Padding(1)
	TitleStyle = lipgloss.NewStyle().Padding(1, 1, 0)

	BoardStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder())

	TileStyle = lipgloss.NewStyle().Width(2).Height(1).AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center)
	FlaggedStyle = TileStyle.Foreground(Lightgray).Background(Black)
	CursorStyle  = TileStyle.Foreground(lipgloss.Color("0")).Background(CursorColor)
	MineStyle    = TileStyle.Background(Black)
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
