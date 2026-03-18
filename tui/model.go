package tui

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/tui/config"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
	"terminal.minesweeper/tui/views"
)

type model struct {
	route    nav.RouteState
	title    views.TitleModel
	sweeper  views.SweeperModel
	settings views.SettingsModel

	width  int
	height int
}

func Model() model {
	return model{
		route:    nav.Title,
		title:    views.MakeTitleModel(),
		sweeper:  views.MakeSweeperModel(),
		settings: views.MakeSettingsModel(),
	}
}

func (m model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case nav.Navigate:
		m.route = msg.Route
		return m, nil
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, config.GameKeyMap.Quit):
			return m, tea.Quit
		}
	}

	switch m.route {
	case nav.Title:
		newTitle, cmd := m.title.Update(msg)
		m.title = newTitle
		return m, cmd
	case nav.Sweeper:
		newSweeper, cmd := m.sweeper.Update(msg)
		m.sweeper = newSweeper
		return m, cmd
	case nav.Settings:
		newSettings, cmd := m.settings.Update(msg)
		m.settings = newSettings
		return m, cmd
	}

	return m, nil
}

func (m model) View() tea.View {
	var v tea.View
	v.AltScreen = true
	var content string

	switch m.route {
	case nav.Title:
		content = m.title.View(m.width, m.height)
	case nav.Sweeper:
		content = m.sweeper.View(m.width, m.height)
	case nav.Settings:
		content = m.settings.View()
	}

	screen := styles.Screen(m.width, m.height).Render(content)
	v.SetContent(screen)

	return v
}
