package views

import (
	"fmt"
	"strconv"
	"time"

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

type tickMsg time.Time

type SweeperModel struct {
	board  *game.Board
	cursor game.Coords
	params *constants.Config
}

func MakeSweeperModel() SweeperModel {
	boardConfig, err := config.LoadConfig()
	if err != nil {
		panic("Read in invalid settings configuration file.")
	}
	return SweeperModel{
		board:  game.GenerateBoard(boardConfig),
		cursor: game.Coords{X: boardConfig.Width / 2, Y: boardConfig.Height / 2},
		params: boardConfig,
	}
}

func Tick() tea.Cmd {
	return tea.Tick(time.Second, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m SweeperModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SweeperModel) Update(msg tea.Msg) (SweeperModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		if m.board.IsStarted() && !m.board.IsComplete() {
			m.board.Tick()
			return m, Tick()
		}

	case nav.Navigate:
		switch msg.Payload {
		case nav.Play:
			m.board = game.GenerateBoard(m.params)
		case nav.Continue:
			return m, Tick()
		}

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
			m.board = game.GenerateBoard(m.params)

		case key.Matches(msg, config.GameKeyMap.Menu):
			if m.board.IsStarted() && !m.board.IsComplete() {
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Title, Payload: nav.Paused}
				}
			}
			return m, func() tea.Msg {
				return nav.Navigate{Route: nav.Title, Payload: nav.New}
			}
		}

		if !m.board.IsComplete() {
			switch {
			case key.Matches(msg, config.GameKeyMap.Flag):
				m.board.SetFlag(m.cursor)
			case key.Matches(msg, config.GameKeyMap.Select):
				if !m.board.IsStarted() {
					m.board.OpenTile(m.cursor)
					return m, Tick()
				}
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
		adjacent := m.board.GetAdjacent(coord)
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

func (m SweeperModel) RenderHeader(width int) string {
	// Render hearts
	hearts := constants.HeartLostSymbol + constants.HeartLostSymbol + constants.HeartSymbol

	lives := styles.Merge([]lipgloss.Style{
		styles.AlignLeft,
		styles.Text(styles.Red),
		styles.Width(width / 2),
	}).Render(hearts)

	// Render flag count
	flagIcon := styles.Highlight(styles.Charcoal).Render(
		fmt.Sprintf("%s ", constants.FlagSymbol),
	)
	flagCount := fmt.Sprintf("%s %s",
		flagIcon,
		styles.Text(styles.Gray).Render(
			fmt.Sprintf("%d/%d",
				m.board.GetFlagCount(),
				m.board.GetMineCount(),
			),
		),
	)

	// Render timer, merge with flag count
	timer := fmt.Sprintf("%s%s", styles.Merge([]lipgloss.Style{
		styles.Highlight(styles.Charcoal),
		styles.PaddingH1,
	}).Render(
		fmt.Sprintf("%05d", m.board.GetTime())),
		styles.Text(styles.Gray).Render("s"),
	)

	right := styles.Merge([]lipgloss.Style{
		styles.AlignRight,
		styles.Width(width / 2),
	}).Render(
		lipgloss.JoinHorizontal(
			lipgloss.Center,
			styles.PaddingH1.Render(flagCount),
			styles.PaddingH1.Render(timer),
		),
	)

	// Merge all
	header := lipgloss.JoinHorizontal(lipgloss.Center, lives, right)
	header = styles.AlignLeft.Width(width).Render(header)

	return header
}

func (m SweeperModel) RenderFooter(width int) string {
	return styles.Merge([]lipgloss.Style{
		styles.AlignHorzCenter,
	}).Width(width - 2).Render(config.RenderHelp(config.GameKeyMap))
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

	board := styles.BoardStyle.Render(
		lipgloss.JoinVertical(
			lipgloss.Center,
			tiles...,
		),
	)

	boardWidth := lipgloss.Width(board)
	boardHeight := lipgloss.Height(board)

	if width < boardWidth || height < boardHeight {
		return styles.AlignCenter.
			Render("Please expand the terminal!")
	}

	header := m.RenderHeader(boardWidth)
	if height < boardHeight+lipgloss.Height(header) {
		return board
	}
	board = lipgloss.JoinVertical(lipgloss.Center, header, board)

	footer := m.RenderFooter(boardWidth)
	if height < boardHeight+lipgloss.Height(footer) {
		return board
	}

	return lipgloss.JoinVertical(lipgloss.Center, board, footer)
}
