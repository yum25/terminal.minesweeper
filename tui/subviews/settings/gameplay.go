package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
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
	return GameplayModel{
		Board:     Board,
		BoardType: BoardType,
		controls:  controls,

		options: [][]fields.Field{
			{
				fields.MakeSelectorModel(
					string(BoardType),
					[]string{
						string(config.BeginnerBoard),
						string(config.IntermediateBoard),
						string(config.AdvancedBoard)},
					controls,
				),
			},
			{},
		},
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
	return "gameplay"
}
