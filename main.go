package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui"
)

func main() {
	p := tea.NewProgram(tui.Model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, looks like the program hit a bomb: %v", err)
		os.Exit(1)
	}
}
