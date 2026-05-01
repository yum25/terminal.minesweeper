package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type AudioModel struct {
	options       []string
	cursor        int
	currentConfig *config.Config
}

func MakeAudioModel(options []string, currentConfig *config.Config) AudioModel {
	return AudioModel{options: options}
}

func (m AudioModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m AudioModel) Update(msg tea.Msg) (AudioModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.currentConfig.UserControls.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.currentConfig.UserControls.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.currentConfig.UserControls.Select):

		}
	}

	return m, nil
}

func (m AudioModel) View(width, height int) string {
	return "audio"
}
