package components

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type SelectorModel struct {
	options  []string
	cursor   int
	controls *config.UserControlsMap
}

func MakeSelectorModel(options []string, controls *config.UserControlsMap) SelectorModel {
	return SelectorModel{options: options}
}

func (m SelectorModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SelectorModel) Update(msg tea.Msg) (SelectorModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.controls.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.controls.Select):

		}
	}

	return m, nil
}

func (m SelectorModel) View(width, height int) string {
	return ""
}
