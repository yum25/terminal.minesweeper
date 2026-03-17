package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type option = string

const (
	configure option = "configure"
	exit      option = "exit"
)

type SettingsModel struct {
	options []option
	cursor  int
}

func MakeSettingsModel() SettingsModel {
	return SettingsModel{
		options: []string{configure, exit},
	}
}

func (m SettingsModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SettingsModel) Update(msg tea.Msg) (SettingsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case "enter", "space":
			switch m.options[m.cursor] {
			case exit:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Title}
				}
			}

		}
	}

	return m, nil
}

func (m SettingsModel) View() string {
	options := make([]string, len(m.options))
	for i, option := range m.options {
		style := styles.OptionStyle
		if i == m.cursor {
			style = styles.SelectedOptionStyle
		}
		options[i] = style.Render(option)
	}

	list := lipgloss.JoinVertical(lipgloss.Center, options...)

	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.IconStyle.Render(constants.MineSymbol),
		styles.TitleStyle.Render("terminal.minesweeper.settings"),
		styles.ListStyle.Render(list),
	)

	return title
}
