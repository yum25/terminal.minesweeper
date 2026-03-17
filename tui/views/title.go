package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type route = string

const (
	play     route = "play"
	settings route = "settings"
	quit     route = "quit"
)

type TitleModel struct {
	paths  []route
	cursor int
}

func MakeTitleModel() TitleModel {
	return TitleModel{
		paths: []route{play, settings, quit},
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
			if m.cursor < len(m.paths)-1 {
				m.cursor++
			}
		case "enter", "space":
			switch m.paths[m.cursor] {
			case play:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Sweeper}
				}
			case settings:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Settings}
				}
			case quit:
				return m, tea.Quit
			}

		}
	}

	return m, nil
}

func (m TitleModel) View() string {
	paths := make([]string, len(m.paths))
	for i, path := range m.paths {
		style := styles.OptionStyle
		if i == m.cursor {
			style = styles.SelectedOptionStyle
		}
		paths[i] = style.Render(path)
	}

	list := lipgloss.JoinVertical(lipgloss.Center, paths...)

	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.IconStyle.Render(constants.MineSymbol),
		styles.TitleStyle.Render("terminal.minesweeper"),
		styles.ListStyle.Render(list),
	)

	return title
}
