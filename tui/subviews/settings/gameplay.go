package settings

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/components/fields"
	"terminal.minesweeper/tui/styles"
)

type selector = string

const (
	PresetSelector selector = "PresetSelector"
	LivesSelector  selector = "LivesSelector"
	WidthSelector  selector = "WidthSelector"
	HeightSelector selector = "HeightSelector"
	MineSelector   selector = "MineSelector"
)

type GameplayModel struct {
	focused bool
	cursorX int
	cursorY int

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
	options := [][]fields.Field{
		{
			fields.MakeSelectorModel(
				PresetSelector,
				&BoardType,
				[]config.BoardPreset{
					config.BeginnerBoard,
					config.IntermediateBoard,
					config.AdvancedBoard},
				controls,
			),
		},
		{
			fields.MakeInputModel(
				LivesSelector,
				&Board.LivesCount,
				"Lives",
				controls,
			),
			fields.MakeInputModel(
				MineSelector,
				&Board.MineCount,
				"Mines",
				controls,
			),
		},
		{
			fields.MakeInputModel(
				WidthSelector,
				&Board.Width,
				"Width",
				controls,
			),
			fields.MakeInputModel(
				HeightSelector,
				&Board.Height,
				"Height",
				controls,
			),
		},
	}

	return GameplayModel{

		Board:     Board,
		BoardType: BoardType,
		controls:  controls,

		options: options,
	}
}

func (m GameplayModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m GameplayModel) Update(msg tea.Msg) (GameplayModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		field := m.options[m.cursorY][m.cursorX]

		if m.focused {
			update, cmd := field.Update(msg)
			m.options[m.cursorY][m.cursorX] = update

			switch {
			case key.Matches(msg, m.controls.Select):
				m.focused = false
			case key.Matches(msg, m.controls.Cancel):
				m.focused = false
			}

			// Side effects
			switch m.BoardType {
			case config.BeginnerBoard:
			case config.IntermediateBoard:
			case config.AdvancedBoard:
			case config.CustomBoard:
			}

			return m, cmd
		}

		switch {
		case key.Matches(msg, m.controls.Up):
			if m.cursorY > 0 {
				m.cursorY--
			}
			m.cursorX = min(len(m.options[m.cursorY])-1, m.cursorX)
		case key.Matches(msg, m.controls.Down):
			if m.cursorY < len(m.options)-1 {
				m.cursorY++
			}
			m.cursorX = min(len(m.options[m.cursorY])-1, m.cursorX)
		case key.Matches(msg, m.controls.Left):
			if m.cursorX > 0 {
				m.cursorX--
			}
		case key.Matches(msg, m.controls.Right):
			if m.cursorX < len(m.options[m.cursorY])-1 {
				m.cursorX++
			}
		case key.Matches(msg, m.controls.Select):
			m.focused = true
		case key.Matches(msg, m.controls.Cancel):
			m.focused = false
		}
	}

	return m, nil
}

func (m GameplayModel) View(width, height int) string {
	containerWidth := width / 2
	fieldHeight := (height / len(m.options)) - 1

	view := make([]string, len(m.options))
	for y, row := range m.options {
		fieldWidth := (containerWidth / len(row)) - 1
		rowView := make([]string, len(row))

		for x, field := range row {
			var hover bool
			if m.cursorY == y && m.cursorX == x {
				hover = true
			} else {
				hover = false
			}

			var state fields.State
			if !hover {
				state = fields.Unfocused
			} else if hover && m.focused {
				state = fields.Focused
			} else {
				state = fields.Hover
			}

			rowView = append(rowView, field.View(fieldWidth, fieldHeight, state))
		}
		view = append(view, lipgloss.JoinHorizontal(lipgloss.Center, rowView...))
	}

	settings := lipgloss.JoinVertical(lipgloss.Center, view...)
	preview := styles.Merge([]lipgloss.Style{
		styles.Width(containerWidth),
		styles.Height(height),
		styles.AlignCenter,
	}).Render("preview")

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render(
		lipgloss.JoinHorizontal(lipgloss.Center, settings, preview),
	)
}
