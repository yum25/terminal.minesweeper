package views

import (
	"slices"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type route = string

const (
	playRoute     route = "play"
	resumeRoute   route = "continue"
	settingsRoute route = "settings"
	quitRoute     route = "quit"
)

type TitleModel struct {
	paths    []route
	cursor   int
	paused   bool
	controls *config.UserControlsMap
}

func MakeTitleModel(controls *config.UserControlsMap) TitleModel {
	return TitleModel{
		paths:    []route{playRoute, settingsRoute, quitRoute},
		controls: controls,
	}
}

func (m TitleModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TitleModel) Update(msg tea.Msg) (TitleModel, tea.Cmd) {
	switch msg := msg.(type) {
	case nav.Navigate:
		switch msg.Payload {
		case nav.New:
			m.paths = []route{playRoute, settingsRoute, quitRoute}
			m.paused = false
		case nav.Paused:
			m.paths = []route{playRoute, resumeRoute, settingsRoute, quitRoute}
			m.cursor = slices.Index(m.paths, resumeRoute)
			m.paused = true
		}

	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Up):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.controls.Down):
			if m.cursor < len(m.paths)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.controls.Select):
			switch m.paths[m.cursor] {
			case playRoute:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Sweeper, Payload: nav.Play}
				}
			case resumeRoute:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Sweeper, Payload: nav.Continue}
				}
			case settingsRoute:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Settings}
				}
			case quitRoute:
				return m, tea.Quit
			}

		case key.Matches(msg, m.controls.Quit):
			return m, tea.Quit
		}

	}

	return m, nil
}

func (m TitleModel) View(width, height int) string {
	paths := []route{playRoute, resumeRoute, settingsRoute, quitRoute}
	for i, path := range paths {
		style := styles.OptionStyle
		if path == m.paths[m.cursor] {
			style = styles.SelectedOptionStyle
		}
		if path == resumeRoute && !m.paused {
			style = styles.DisabledOptionStyle
		}
		paths[i] = style.Render(path)
	}

	list := lipgloss.JoinVertical(lipgloss.Center, paths...)
	title := lipgloss.JoinVertical(lipgloss.Center,
		styles.IconStyle.Render(constants.MineSymbol),
		styles.TitleStyle.Render("terminal.minesweeper"),
		styles.ListStyle.Render(list),
	)

	footer := styles.Merge([]lipgloss.Style{
		styles.AlignBottom,
		styles.AlignHorzCenter,
		styles.Width(width),
	}).Render(config.RenderHelp(m.controls))

	styles.Merge([]lipgloss.Style{
		styles.AlignCenter,
		styles.Width(width),
		styles.Height(height - lipgloss.Height(footer)),
	})

	title = styles.Merge([]lipgloss.Style{
		styles.AlignCenter,
		styles.Width(width),
		styles.Height(height - lipgloss.Height(footer)),
	}).Render(title)

	return lipgloss.JoinVertical(lipgloss.Center, title, footer)
}
