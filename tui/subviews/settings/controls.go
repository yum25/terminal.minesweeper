package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type ControlsModel struct {
	options  []string
	cursor   int
	controls *config.UserControlsMap
}

func MakeControlsModel(controls *config.GameControlsMap) ControlsModel {
	return ControlsModel{controls: &controls.UserControlsMap}
}

func (m ControlsModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m ControlsModel) Update(msg tea.Msg, bindMode bool) (ControlsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Up):
		case key.Matches(msg, m.controls.Down):
		case key.Matches(msg, m.controls.Select):
		}
	}

	return m, nil
}

func (m ControlsModel) View(width, height int, bindMode bool) string {
	return "controls"
}
