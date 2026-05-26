package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/styles"
)

type AudioModel struct {
	cursor   int
	controls *config.UserControlsMap
}

func MakeAudioModel(controls *config.UserControlsMap) AudioModel {
	return AudioModel{controls: controls}
}

func (m AudioModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m AudioModel) Update(msg tea.Msg) (AudioModel, tea.Cmd) {
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

func (m AudioModel) View(width, height int) string {
	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render("Coming soon!")
}
