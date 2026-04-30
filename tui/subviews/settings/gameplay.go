package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
)

type GameplayModel struct {
	options []string
	cursor  int
}

func MakeGameplayModel(options []string) GameplayModel {
	return GameplayModel{options: options}
}

func (m GameplayModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GameplayModel) Update(msg tea.Msg) (GameplayModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, config.UserKeyMap.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, config.UserKeyMap.Down):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, config.UserKeyMap.Select):

		}
	}

	return m, nil
}

func (m GameplayModel) View(width, height int) string {
	return "gameplay"
}
