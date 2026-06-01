package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/components/fields"
	"terminal.minesweeper/tui/styles"
)

type ControlsModel struct {
	cursorX int
	cursorY int

	preset       *config.ControlPreset
	keys         *config.Controls
	controls     *config.UserControlsMap
	gameControls *config.GameControlsMap

	options [][]fields.Field
}

func MakeControlsModel(c *config.Config) ControlsModel {
	keys := c.FromKeyMap()
	options := [][]fields.Field{
		{
			fields.MakeSelectorModel(
				"",
				&c.ControlType,
				[]config.ControlPreset{
					config.DefaultControls,
					config.VimControls,
					config.Custom,
				},
				&c.UserControls,
			),
		},
		{
			fields.MakeInputModel(
				"",
				&keys.Up,
				"Up",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Down,
				"Down",
				&c.UserControls,
			),
		},
		{
			fields.MakeInputModel(
				"",
				&keys.Left,
				"Left",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Right,
				"Right",
				&c.UserControls,
			),
		},
		{
			fields.MakeInputModel(
				"",
				&keys.Select,
				"Select",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Menu,
				"Menu",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Cancel,
				"Cancel",
				&c.UserControls,
			),
		},
		{
			fields.MakeInputModel(
				"",
				&keys.Quit,
				"Quit",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Flag,
				"Flag",
				&c.UserControls,
			),
			fields.MakeInputModel(
				"",
				&keys.Restart,
				"Restart",
				&c.UserControls,
			),
		},
	}

	return ControlsModel{
		preset:       &c.ControlType,
		keys:         &keys,
		controls:     &c.UserControls,
		gameControls: &c.GameControls,
		options:      options,
	}
}

func (m ControlsModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m ControlsModel) Update(msg tea.Msg, bindMode bool) (ControlsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		field := m.options[m.cursorY][m.cursorX]

		if bindMode {
			update, cmd := field.Update(msg)
			m.options[m.cursorY][m.cursorX] = update

			switch *m.preset {
			case config.DefaultControls:
				*m.keys = config.DEFAULT_CONTROLS
			case config.VimControls:
				*m.keys = config.VIM_CONTROLS
			case config.Custom:
			}

			*m.controls, *m.gameControls = m.keys.ToKeyMap()
			return m, cmd
		}

		switch {
		case key.Matches(msg, m.controls.Up):
			if m.cursorY > 0 {
				m.cursorY--
			}
			m.cursorX = min(len(m.options[m.cursorY])-1, m.cursorX)
		case key.Matches(msg, m.controls.Down):
			if m.cursorY < len(m.options)-1 {
				m.cursorY++
			}
			m.cursorX = min(len(m.options[m.cursorY])-1, m.cursorX)
		case key.Matches(msg, m.controls.Left):
			if m.cursorX > 0 {
				m.cursorX--
			}
		case key.Matches(msg, m.controls.Right):
			if m.cursorX < len(m.options[m.cursorY])-1 {
				m.cursorX++
			}
		}
	}

	return m, nil
}

func (m ControlsModel) View(width, height int, bindMode bool) string {
	fieldHeight := (height / len(m.options))

	view := make([]string, len(m.options))
	for y, row := range m.options {
		fieldWidth := (width / 3)
		rowView := make([]string, len(row))

		for x, field := range row {
			var hover bool
			if m.cursorY == y && m.cursorX == x {
				hover = true
			} else {
				hover = false
			}

			var state fields.State
			if !hover {
				state = fields.Unfocused
			} else if hover && bindMode {
				state = fields.Focused
			} else {
				state = fields.Hover
			}

			rowView[x] = field.View(fieldWidth, fieldHeight, state)
		}
		view[y] = styles.AlignCenter.Render(lipgloss.JoinHorizontal(lipgloss.Center, rowView...))
	}

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render(lipgloss.JoinVertical(lipgloss.Center, view...))
}
