package views

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/game"
)

type SweeperModel struct {
	board  *game.Board
	cursor game.Coords
}

func MakeSweeperModel() SweeperModel {
	return SweeperModel{
		board:  game.GenerateBoard(game.Width, game.Height, game.MineCount),
		cursor: game.Coords{X: 12, Y: 10},
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
			if m.cursor.Y > 0 {
				m.cursor.Y--
			}

		case "down", "s":
			if m.cursor.Y < m.board.GetHeight()-1 {
				m.cursor.Y++
			}

		case "left", "a":
			if m.cursor.X > 0 {
				m.cursor.X--
			}

		case "right", "d":
			if m.cursor.X < m.board.GetWidth()-1 {
				m.cursor.X++
			}

		case "f":
			m.board.Flag(m.cursor)

		case "enter", "space":
			m.board.Reveal(m.cursor)
		}
	}

	return m, nil
}

func (m SweeperModel) RenderTile(coord game.Coords) string {
	style := tileStyle
	var tileContent string

	if m.board.IsFlagged(coord) {
		tileContent = flagSymbol
		style = FlaggedStyle(m.board.IsComplete(), m.board.IsMine(coord))
	} else if m.board.IsRevealed(coord) {
		adjacent := m.board.Adjacent(coord)
		tileContent = strconv.Itoa(adjacent)
		style = TileStyle(adjacent)
	} else {
	}

	if coord.X == m.cursor.X && coord.Y == m.cursor.Y {
		tileContent += cursorSymbol
		style = cursor
	}

	return style.Render(tileContent)
}

func (m SweeperModel) View() string {
	tiles := make([]string, m.board.GetHeight())

	for y := range m.board.GetHeight() {
		row := make([]string, m.board.GetWidth())
		for x := range row {
			row[x] = m.RenderTile(game.Coords{X: x, Y: y})
		}

		tiles[y] = lipgloss.JoinHorizontal(lipgloss.Center, row...)
	}

	board := boardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			tiles...,
		),
	)

	return board
}
