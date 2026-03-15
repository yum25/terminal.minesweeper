package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type Option = string

type TitleModel struct {
	options []Option
	cursor  int
}

func MakeTitleModel() TitleModel {
	return TitleModel{
		options: []string{"play", "quit"},
	}
}

func (m TitleModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TitleModel) Update(msg tea.Msg) (TitleModel, tea.Cmd) {
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
			case "play":
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Sweeper}
				}
			case "quit":
				return m, tea.Quit
			}

		}
	}

	return m, nil
}

func (m TitleModel) View() string {
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
		styles.TitleStyle.Render("terminal.minesweeper"),
		styles.ListStyle.Render(list),
	)

	return title
}
