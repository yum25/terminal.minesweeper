package views

import (
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Option = string

type TitleModel struct {
	options []Option
	cursor  int

	width  int
	height int
}

func MakeTitleModel() TitleModel {
	return TitleModel{
		options: []string{"play", "quit"},
	}
}

func (m TitleModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m TitleModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyPressMsg:
		switch msg.String() {

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case "enter", "space":
			switch m.options[m.cursor] {
			case "play":
			case "quit":
				return m, tea.Quit
			}

		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m TitleModel) View() tea.View {
	var v tea.View
	v.AltScreen = true

	options := make([]string, len(m.options))
	for i, option := range m.options {
		style := optionStyle
		if i == m.cursor {
			style = selectedOptionStyle
		}
		options[i] = style.Render(option)
	}

	list := lipgloss.JoinVertical(lipgloss.Center, options...)

	title := lipgloss.JoinVertical(lipgloss.Center,
		iconStyle.Render("☀"),
		titleStyle.Render("terminal.minesweeper"),
		listStyle.Render(list),
	)

	content := Screen(m.width, m.height).Render(title)
	v.SetContent(content)
	return v
}
