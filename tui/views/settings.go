package views

import (
	"slices"
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
	save     option = "save"
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
		options:       []string{gameplay, display, audio, controls},
		focus:         gameplay,
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
		case key.Matches(msg, m.localConfig.UserControls.Menu):
			return m, func() tea.Msg {
				return nav.Navigate{Route: nav.Title}
			}
		case key.Matches(msg, m.localConfig.UserControls.Up) ||
			key.Matches(msg, m.localConfig.UserControls.Left):
			// Pass into focused
		case key.Matches(msg, m.localConfig.UserControls.Down) ||
			key.Matches(msg, m.localConfig.UserControls.Right):
			// Pass into focused
		case key.Matches(msg, m.localConfig.UserControls.Select):
			// Pass into focused
		default:
			if len(msg.String()) == 1 && msg.String() >= "1" && msg.String() <= strconv.Itoa(len(m.options)) {
				num, _ := strconv.Atoi(msg.String())

				m.cursor = num - 1
				m.focus = m.options[m.cursor]
			}
		}
	}

	return m, nil
}

func (m SettingsModel) RenderOption(o option) string {
	style := styles.OptionStyle
	if o == m.options[m.cursor] {
		style = styles.SelectedOptionStyle
	}

	var key string
	switch index := slices.Index(m.options, o); index {
	case -1:
		switch o {
		case save:
			key = styles.IndentStyle.Render("ctrl+s")
		case exit:
			key = styles.IndentStyle.Render(m.localConfig.UserControls.Menu.Keys()...)
		}
	default:
		key = styles.IndentStyle.Render(strconv.Itoa(index + 1))
	}

	card := styles.Merge([]lipgloss.Style{
		style,
		styles.Width(
			lipgloss.Width(o) + lipgloss.Width(key) + 3),
	}).Render(
		lipgloss.JoinHorizontal(lipgloss.Center,
			o,
			" ",
			key,
		),
	)

	if o == m.options[m.cursor] {
		card = styles.AddHalfPixelBorder(card, styles.Text(styles.CursorColor))
	}
	card = styles.PaddingH1.Render(card)
	return card
}

func (m SettingsModel) View(width, height int) string {
	options := make([]string, len(m.options))
	for i, o := range m.options {
		options[i] = m.RenderOption(o)
	}

	list := lipgloss.JoinHorizontal(lipgloss.Center, options...)

	footer := lipgloss.JoinHorizontal(lipgloss.Center,
		m.RenderOption(save),
		m.RenderOption(exit),
	)

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
		footer,
	)

	return title
}
