package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui"
)

const dev = true

func main() {
	if dev {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()

	}

	p := tea.NewProgram(tui.Model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, looks like the program hit a bomb: %v", err)
		os.Exit(1)
	}
}
