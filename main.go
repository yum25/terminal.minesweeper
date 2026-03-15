package main

import (
	"fmt"
	"log"
	"os"

	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui"
)

const dev = false

func main() {
	if dev {
		f, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		log.SetOutput(f)
	}

	p := tea.NewProgram(tui.Model())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Oops, looks like the program hit a bomb: %v", err)
		os.Exit(1)
	}
}
