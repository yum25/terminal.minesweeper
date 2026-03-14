package tui

import (
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui/views"
)

type routeState int

const (
	title routeState = iota
	sweeper
)

type model struct {
	route   routeState
	title   views.TitleModel
	sweeper views.SweeperModel
}

func Model() model {
	return model{
		route:   title,
		title:   views.MakeTitleModel(),
		sweeper: views.MakeSweeperModel(),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}

	}

	switch m.route {
	case title:
		newTitle, cmd := m.title.Update(msg)
		m.title = newTitle.(views.TitleModel)
		return m, cmd
	case sweeper:
		newSweeper, cmd := m.title.Update(msg)
		m.title = newSweeper.(views.TitleModel)
		return m, cmd
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m model) View() tea.View {
	switch m.route {

	case title:
		return m.title.View()

	case sweeper:

	}

	return tea.NewView("")
}
