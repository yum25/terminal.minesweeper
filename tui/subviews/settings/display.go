package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/styles"
)

type DisplayModel struct {
	cursor   int
	controls *config.UserControlsMap
}

func MakeDisplayModel(controls *config.UserControlsMap) DisplayModel {
	return DisplayModel{controls: controls}
}

func (m DisplayModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m DisplayModel) Update(msg tea.Msg, bindMode bool) (DisplayModel, tea.Cmd) {
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

func (m DisplayModel) View(width, height int, bindMode bool) string {

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render("Coming soon!")
}
