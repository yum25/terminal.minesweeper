package views

import (
	"strconv"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/nav"
	"terminal.minesweeper/tui/styles"
	"terminal.minesweeper/tui/subviews/settings"
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
	options       []option
	cursor        int
	focus         option
	bindMode      bool
	localConfig   *config.Config
	currentConfig *config.Config
	// Subviews
	GameplayView settings.GameplayModel
	DisplayView  settings.DisplayModel
	AudioView    settings.AudioModel
	ControlsView settings.ControlsModel
}

func MakeSettingsModel(currentConfig *config.Config) SettingsModel {
	var localConfig config.Config
	err := localConfig.LoadConfig()

	if err != nil {
		localConfig = config.DEFAULT_CONFIG
	}
	return SettingsModel{
		options:       []string{gameplay, display, audio, controls, exit},
		localConfig:   &localConfig,
		currentConfig: currentConfig,
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
		case key.Matches(msg, m.localConfig.UserControls.Up) ||
			key.Matches(msg, m.localConfig.UserControls.Left):
			if m.cursor > 0 {
				m.cursor--
			}
		case key.Matches(msg, m.localConfig.UserControls.Down) ||
			key.Matches(msg, m.localConfig.UserControls.Right):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}
		case key.Matches(msg, m.localConfig.UserControls.Select):
			switch m.options[m.cursor] {
			case exit:
				return m, func() tea.Msg {
					return nav.Navigate{Route: nav.Title}
				}
			default:
				m.focus = m.options[m.cursor]
			}
		default:
			if len(msg.String()) == 1 && msg.String() >= "0" && msg.String() <= "5" {
				num, _ := strconv.Atoi(msg.String())

				m.cursor = num - 1
				m.focus = m.options[m.cursor]
			}
		}
	}

	return m, nil
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

	var view string
	switch m.focus {
	case gameplay:
		view = m.GameplayView.View(width, height)
	case display:
		view = m.DisplayView.View(width, height)
	case audio:
		view = m.AudioView.View(width, height)
	case controls:
		view = m.ControlsView.View(width, height)
	}

	title := lipgloss.JoinVertical(lipgloss.Center,
		list,
		container.Render(view),
	)

	return title
}
