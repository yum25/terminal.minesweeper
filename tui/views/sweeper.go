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
}

func MakeSweeperModel() SweeperModel {
	return SweeperModel{
		board:   game.GenerateBoard(24, 20, 99),
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

	case tea.KeyPressMsg:
		switch msg.String() {

		case "up", "w":
			if m.cursorY > 0 {
				m.cursorY--
			}

		case "down", "s":
			if m.cursorY < m.board.GetHeight()-1 {
				m.cursorY++
			}

		case "left", "a":
			if m.cursorX > 0 {
				m.cursorX--
			}

		case "right", "d":
			if m.cursorX < m.board.GetWidth()-1 {
				m.cursorX++
			}

		case "f":
			m.board.Flag(m.cursorY, m.cursorX)

		case "enter", "space":

		}
	}

	return m, nil
}

func (m SweeperModel) View() string {
	columns := make([]string, m.board.GetHeight())

	for y := range m.board.GetHeight() {
		tiles := make([]string, m.board.GetWidth())

		for x := range tiles {
			style := tileStyle
			if x == m.cursorX && y == m.cursorY {
				style = cursor
			}

			if m.board.IsFlagged(y, x) {
				tiles[x] = style.Render("⚑")
			} else {
				tiles[x] = style.Render("")
			}
			tiles[x] = style.Render("")

		}
		row := lipgloss.JoinHorizontal(lipgloss.Center, tiles...)

		columns[y] = row
	}

	board := boardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			columns...,
		),
	)

	return board
}
