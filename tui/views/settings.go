package views

import (
	"strconv"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/tui/config"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
)

type option = string

const (
	gameplay option = "gameplay"
	display  option = "display"
	audio    option = "audio"
	controls option = "controls"
	exit     option = "exit"
)

type SettingsModel struct {
	options  []option
	cursor   int
	focus    option
	bindMode bool
}

func MakeSettingsModel() SettingsModel {
	return SettingsModel{
		options: []string{gameplay, display, audio, controls, exit},
	}
}

func (m SettingsModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SettingsModel) Update(msg tea.Msg) (SettingsModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, config.GameKeyMap.Up) ||
			key.Matches(msg, config.GameKeyMap.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, config.GameKeyMap.Down) ||
			key.Matches(msg, config.GameKeyMap.Right):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, config.GameKeyMap.Select):
			switch m.options[m.cursor] {
			case exit:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Title}
				}
			}
		}
	}

	return m, nil
}

func (m SettingsModel) RenderGameplay() string {
	return ""
}

func (m SettingsModel) RenderDisplay() string {
	return ""
}

func (m SettingsModel) RenderAudio() string {
	return ""
}

func (m SettingsModel) RenderControls() string {
	return ""
}

func (m SettingsModel) View(width, height int) string {
	options := make([]string, len(m.options))
	for i, option := range m.options {
		style := styles.OptionStyle
		if i == m.cursor {
			style = styles.SelectedOptionStyle
		}

		index := styles.IndentStyle.Render(strconv.Itoa(i + 1))
		options[i] = styles.Merge([]lipgloss.Style{
			style,
			styles.Width(
				lipgloss.Width(option) + lipgloss.Width(index) + 3),
		}).Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				option,
				" ",
				index,
			),
		)

		if i == m.cursor {
			options[i] = styles.AddHalfPixelBorder(options[i], styles.Text(styles.CursorColor))
		}

		options[i] = styles.PaddingH1.Render(options[i])
	}

	list := lipgloss.JoinHorizontal(lipgloss.Center, options...)

	container := styles.Merge([]lipgloss.Style{
		styles.BorderStyle,
		styles.Width(min(width-6, 100)),
		styles.Height(min(height-lipgloss.Height(list), 25)),
	})

	title := lipgloss.JoinVertical(lipgloss.Center,
		list,
		container.Render(""),
	)

	return title
}
