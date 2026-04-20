package main

import (
	"fmt"
	"os"

	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui"
	"terminal.minesweeper/tui/config"
)

func main() {
	if len(os.Getenv("DEBUG")) > 0 {
		f, err := tea.LogToFile("debug.log", "debug")
		if err != nil {
			fmt.Println("fatal:", err)
			os.Exit(1)
		}
		defer f.Close()
	}

	err := config.LoadConfig()
	if err != nil {
		config.Current = config.DEFAULT_CONFIG
	}

	p := tea.NewProgram(tui.Model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, looks like the program hit a bomb: %v", err)
		os.Exit(1)
	}
}
