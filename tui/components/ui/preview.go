package ui

import (
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/styles"
)

func RenderPreview(Board *config.BoardConfig) string {
	rows := []string{}
	for y := 0; y < Board.Height; y += 2 {
		row := []string{}
		for x := 0; x < Board.Width; x += 1 {
			bottom := y + 1

			var cell string
			if bottom < Board.Height {
				cell = styles.Merge([]lipgloss.Style{
					styles.Text(styles.Green),
					styles.Highlight(styles.Green),
				},
				).Render(constants.HalfPixelTop)
			} else {
				cell = styles.Text(styles.Green).Render(constants.HalfPixelTop)
			}

			row = append(row, cell)
		}
		rows = append(rows, lipgloss.JoinHorizontal(lipgloss.Center, row...))
	}

	return lipgloss.JoinVertical(lipgloss.Center, rows...)
}
