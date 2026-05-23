package fields

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/styles"
)

type SelectorModel struct {
	value    string
	options  []string
	cursor   int
	controls *config.UserControlsMap
}

func MakeSelectorModel(value string, options []string, controls *config.UserControlsMap) SelectorModel {
	return SelectorModel{value: value, options: options}
}

func (m SelectorModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SelectorModel) Update(msg tea.Msg) (Field, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.controls.Right):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case key.Matches(msg, m.controls.Select):
			m.value = m.options[m.cursor]
		}
	}

	return m, nil
}

func (m SelectorModel) View(width, height int, focused bool) string {
	option := styles.OptionStyle.Render(m.value)
	if focused {
		option = styles.SelectedOptionStyle.Render(m.options[m.cursor])
	}

	return lipgloss.JoinHorizontal(lipgloss.Center,
		constants.ArrowLeftSymbol,
		option,
		constants.ArrowRightSymbol)
}
