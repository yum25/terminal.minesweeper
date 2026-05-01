package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type GameplayModel struct {
	cursor    int
	options   []string
	Board     *config.BoardConfig
	BoardType *config.BoardPreset
	controls  *config.UserControlsMap
}

func MakeGameplayModel(
	Board *config.BoardConfig,
	BoardType *config.BoardPreset,
	controls *config.UserControlsMap,
) GameplayModel {
	return GameplayModel{Board: Board, BoardType: BoardType, controls: controls}
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

func (m GameplayModel) View(width, height int) string {
	return "gameplay"
}
