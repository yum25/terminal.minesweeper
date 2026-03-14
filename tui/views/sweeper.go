package views

import (
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui/game"
)

type SweeperModel struct {
	board  game.Board
	cursor []int

	width  int
	height int
}

func MakeSweeperModel() SweeperModel {
	return SweeperModel{}
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

		case "up", "k":

		case "down", "j":

		case "enter", "space":

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m SweeperModel) View() string {
	title := titleStyle.Render("terminal.minesweeper.board")

	return title
}
