package components

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type InputModel struct {
	options  []string
	cursor   int
	controls *config.UserControlsMap
}

func MakeInputModel(options []string, controls *config.UserControlsMap) InputModel {
	return InputModel{options: options, controls: controls}
}

func (m InputModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
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

func (m InputModel) View(width, height int) string {
	return ""
}
