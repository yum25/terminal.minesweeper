package views

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/config"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type option = string

const (
	gameplay option = "gameplay"
	display  option = "display"
	audio    option = "audio"
	controls option = "controls"
	exit     option = "exit"
)

type SettingsModel struct {
	options  []option
	cursor   int
	focus    option
	bindMode bool
}

func MakeSettingsModel() SettingsModel {
	return SettingsModel{
		options: []string{gameplay, display, audio, controls, exit},
	}
}

func (m SettingsModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SettingsModel) Update(msg tea.Msg) (SettingsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, config.GameKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, config.GameKeyMap.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, config.GameKeyMap.Select):
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

	list := lipgloss.JoinHorizontal(lipgloss.Center, options...)

	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.ListStyle.Render(list),
		styles.BoardStyle.Render(""),
	)

	return title
}
