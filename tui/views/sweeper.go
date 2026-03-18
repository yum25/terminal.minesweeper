package views

import (
	"strconv"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/game"
	state "terminal.minesweeper/shared"
	"terminal.minesweeper/tui/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/nav"
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
		switch {
		case key.Matches(msg, config.GameKeyMap.Up):
			if m.cursor.Y > 0 {
				m.cursor.Y--
			}
		case key.Matches(msg, config.GameKeyMap.Down):
			if m.cursor.Y < m.board.GetHeight()-1 {
				m.cursor.Y++
			}
		case key.Matches(msg, config.GameKeyMap.Left):
			if m.cursor.X > 0 {
				m.cursor.X--
			}
		case key.Matches(msg, config.GameKeyMap.Right):
			if m.cursor.X < m.board.GetWidth()-1 {
				m.cursor.X++
			}
		case key.Matches(msg, config.GameKeyMap.Restart):
			m.board = game.GenerateBoard(config.Width, config.Height, config.MineCount)
		case key.Matches(msg, config.GameKeyMap.Menu):
			return m, func() tea.Msg {
				return nav.Navigate{Route: nav.Title}
			}
		}
		if !m.board.IsComplete() {
			switch {
			case key.Matches(msg, config.GameKeyMap.Flag):
				m.board.Flag(m.cursor)
			case key.Matches(msg, config.GameKeyMap.Select):
				m.board.OpenTile(m.cursor)
			}

			m.board.CheckIsComplete()
		}
	}

	return m, nil
}

func (m SweeperModel) RenderTile(coord game.Coords) string {
	style := styles.TileStyle
	var tileContent string

	switch m.board.GetTileState(coord) {
	case state.TileFlagged:
		tileContent = constants.FlagSymbol
		style = styles.FlaggedStyle
	case state.TileFlaggedWrong:
		tileContent = constants.FlagSymbol
		style = styles.FlaggedStyle
		if m.board.IsComplete() {
			style = styles.MineStyle
		}
	case state.TileOpen:
		adjacent := m.board.Adjacent(coord)
		tileContent = strconv.Itoa(adjacent)
		style = styles.RevealedStyle(adjacent)
	case state.MineHit:
		tileContent = constants.MineHitSymbol
		style = styles.MineHitStyle
	case state.MineRevealed:
		tileContent = constants.MineSymbol
		style = styles.MineStyle
	}

	if coord.X == m.cursor.X && coord.Y == m.cursor.Y {
		tileContent += constants.CursorSymbol
		style = styles.CursorStyle
	}

	return style.Render(tileContent)
}

func (m SweeperModel) View(width, height int) string {
	tiles := make([]string, m.board.GetHeight())

	for y := range m.board.GetHeight() {
		row := make([]string, m.board.GetWidth())
		for x := range row {
			row[x] = m.RenderTile(game.Coords{X: x, Y: y})
		}

		tiles[y] = lipgloss.JoinHorizontal(lipgloss.Center, row...)
	}

	footer := lipgloss.NewStyle().Padding(1, 2).Render(config.RenderHelp(config.GameKeyMap))
	board := styles.BoardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			tiles...,
		),
	)

	boardWidth := lipgloss.Width(board)
	boardHeight := lipgloss.Height(board)

	if width < boardWidth || height < boardHeight {
		return lipgloss.NewStyle().AlignHorizontal(lipgloss.Center).AlignVertical(lipgloss.Center).
			Render("Please expand the terminal!")
	}

	if height < boardHeight+lipgloss.Height(footer) {
		return board
	}

	board = lipgloss.NewStyle().Width(width).AlignHorizontal(lipgloss.Center).Render(board)
	footer = lipgloss.NewStyle().Width(width).AlignHorizontal(lipgloss.Center).Render(footer)

	return lipgloss.JoinVertical(lipgloss.Center, board, footer)
}
