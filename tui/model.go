package tui

import (
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui/views"
)

type model struct {
	route   views.RouteState
	title   views.TitleModel
	sweeper views.SweeperModel

	width  int
	height int
}

func Model() model {
	return model{
		route:   views.Title,
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
	case views.Navigate:
		m.route = views.Sweeper
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	switch m.route {
	case views.Title:
		newTitle, cmd := m.title.Update(msg)
		m.title = newTitle
		return m, cmd
	case views.Sweeper:
		newSweeper, cmd := m.title.Update(msg)
		m.title = newSweeper
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	var v tea.View
	v.AltScreen = true

	var content string

	switch m.route {

	case views.Title:
		content = m.title.View()
	case views.Sweeper:
		content = m.sweeper.View()
	}

	screen := views.Screen(m.width, m.height).Render(content)
	v.SetContent(screen)

	return v
}
