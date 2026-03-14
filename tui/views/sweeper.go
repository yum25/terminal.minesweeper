package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/game"
)

type SweeperModel struct {
	board   *game.Board
	cursorX int
	cursorY int

	width  int
	height int
}

func MakeSweeperModel() SweeperModel {
	return SweeperModel{
		board:   game.GenerateBoard(24, 20, 99),
		width:   24,
		height:  20,
		cursorX: 12,
		cursorY: 10,
	}
}

func (m SweeperModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SweeperModel) Update(msg tea.Msg) (SweeperModel, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {

		case "up", "w":
			if m.cursorY > 0 {
				m.cursorY--
			}

		case "down", "s":
			if m.cursorY < m.height-1 {
				m.cursorY++
			}

		case "left", "a":
			if m.cursorX > 0 {
				m.cursorX--
			}

		case "right", "d":
			if m.cursorX < m.width-1 {
				m.cursorX++
			}

		case "enter", "space":

		}
	}

	return m, nil
}

func (m SweeperModel) View() string {
	columns := make([]string, m.height)

	for i := range m.height {
		tiles := make([]string, m.width)
		for j := range tiles {
			tiles[j] = tile

			if j == m.cursorX && i == m.cursorY {
				tiles[j] = cursor
			}
		}
		row := lipgloss.JoinHorizontal(lipgloss.Center, tiles...)

		columns[i] = row
	}

	board := boardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			columns...,
		),
	)

	return board
}
