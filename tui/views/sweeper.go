package views

import (
	"strconv"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/game"
	state "terminal.minesweeper/shared"
	"terminal.minesweeper/tui/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/styles"
)

type SweeperModel struct {
	board  *game.Board
	cursor game.Coords
}

func MakeSweeperModel() SweeperModel {
	return SweeperModel{
		board:  game.GenerateBoard(config.Width, config.Height, config.MineCount),
		cursor: game.Coords{X: config.Width / 2, Y: config.Height / 2},
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
			m.board.OpenTile(m.cursor)
		}
	}

	return m, nil
}

func (m SweeperModel) RenderTile(coord game.Coords) string {
	style := styles.TileStyle
	var tileContent string

	switch m.board.GetTileState(coord) {
	// General states
	case state.TileFlagged:
		tileContent = constants.FlagSymbol
		style = styles.FlaggedStyle
	case state.TileFlaggedWrong:
		tileContent = constants.FlagSymbol
		style = styles.FlaggedStyle
	case state.TileOpen:
		adjacent := m.board.Adjacent(coord)
		tileContent = strconv.Itoa(adjacent)
		style = styles.RevealedStyle(adjacent)
	case state.MineHit:
		tileContent = constants.MineSymbol
		style = styles.MineStyle
	}

	if coord.X == m.cursor.X && coord.Y == m.cursor.Y {
		tileContent += constants.CursorSymbol
		style = styles.CursorStyle
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

	board := styles.BoardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			tiles...,
		),
	)

	return board
}
