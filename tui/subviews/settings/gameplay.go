package settings

import (
	"slices"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/components/fields"
)

type GameplayModel struct {
	focused bool
	cursorX int
	cursorY int
	rows    int
	columns int

	Board     *config.BoardConfig
	BoardType config.BoardPreset
	controls  *config.UserControlsMap

	options [][]fields.Field
}

func MakeGameplayModel(
	Board *config.BoardConfig,
	BoardType config.BoardPreset,
	controls *config.UserControlsMap,
) GameplayModel {
	options := [][]fields.Field{
		{
			fields.MakeSelectorModel(
				&BoardType,
				[]config.BoardPreset{
					config.BeginnerBoard,
					config.IntermediateBoard,
					config.AdvancedBoard},
				controls,
			),
		},
	}

	colLens := make([]int, len(options))
	for _, c := range options {
		colLens = append(colLens, len(c))
	}

	return GameplayModel{
		rows:    max(len(options), 1),
		columns: max(slices.Max(colLens), 1),

		Board:     Board,
		BoardType: BoardType,
		controls:  controls,

		options: options,
	}
}

func (m GameplayModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GameplayModel) Update(msg tea.Msg) (GameplayModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Up):
			if m.cursorY > 0 {
				m.cursorY--
			}
		case key.Matches(msg, m.controls.Down):
			if m.cursorY < m.rows-1 {
				m.cursorY++
			}
		case key.Matches(msg, m.controls.Left):
			if m.cursorX > 0 {
				m.cursorX--
			}
		case key.Matches(msg, m.controls.Right):
			if m.cursorX < m.columns-1 {
				m.cursorX++
			}
		case key.Matches(msg, m.controls.Select):
			m.focused = !m.focused
		case key.Matches(msg, m.controls.Cancel):
			m.focused = false
		}
	}

	return m, nil
}

func (m GameplayModel) View(width, height int) string {
	fieldWidth := width / m.columns
	fieldHeight := height / m.rows

	view := make([]string, len(m.options))
	for y, row := range m.options {
		rowView := make([]string, len(row))
		for x, field := range row {
			var hover bool
			if m.cursorY == y && m.cursorX == x {
				hover = true
			} else {
				hover = false
			}

			rowView = append(rowView, field.View(fieldWidth, fieldHeight, hover && m.focused))
		}
		view = append(view, lipgloss.JoinHorizontal(lipgloss.Center, rowView...))
	}

	return lipgloss.JoinVertical(lipgloss.Center, view...)
}
