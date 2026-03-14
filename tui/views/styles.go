package views

import lipgloss "charm.land/lipgloss/v2"

var (
	optionStyle = lipgloss.NewStyle().
			Padding(0, 2)

	selectedOptionStyle = lipgloss.NewStyle().
				Padding(0, 2).
				Foreground(lipgloss.Color("0")).
				Background(lipgloss.Color("#FF5733")).
				Bold(true)

	iconStyle = lipgloss.NewStyle().
			Background(lipgloss.Color("#9AC085")).
			Padding(0, 1)

	listStyle  = lipgloss.NewStyle().Padding(1)
	titleStyle = lipgloss.NewStyle().Padding(0, 1)

	tileStyle = lipgloss.NewStyle().Padding(0, 0).Width(2).Height(1).AlignHorizontal(lipgloss.Center).
			AlignVertical(lipgloss.Center)
	boardStyle = lipgloss.NewStyle().Padding(0, 0).Border(lipgloss.RoundedBorder())

	cursor = tileStyle.Background(lipgloss.Color("#FF5733"))
)

func Screen(width int, height int) lipgloss.Style {
	return lipgloss.NewStyle().
		Width(width).
		Height(height).
		AlignHorizontal(lipgloss.Center).
		AlignVertical(lipgloss.Center)
}
