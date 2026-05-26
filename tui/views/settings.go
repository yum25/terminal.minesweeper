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

		GameplayView: settings.MakeGameplayModel(
			&localConfig.Board,
			localConfig.BoardType,
			&localConfig.UserControls,
		),
		DisplayView: settings.MakeDisplayModel(
			&localConfig.UserControls,
		),
		AudioView: settings.MakeAudioModel(

			&localConfig.UserControls,
		),
		ControlsView: settings.MakeControlsModel(
			&localConfig.GameControls,
		),
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
		case len(msg.String()) == 1 && msg.String() >= "1" && msg.String() <= strconv.Itoa(len(m.options)):
			num, _ := strconv.Atoi(msg.String())

			m.cursor = num - 1
			m.focus = m.options[m.cursor]
		default:
			switch m.focus {
			case gameplay:
				view, cmd := m.GameplayView.Update(msg)
				m.GameplayView = view
				return m, cmd
			case display:
				view, cmd := m.DisplayView.Update(msg)
				m.DisplayView = view
				return m, cmd
			case audio:
				view, cmd := m.AudioView.Update(msg)
				m.AudioView = view
				return m, cmd
			case controls:
				view, cmd := m.ControlsView.Update(msg)
				m.ControlsView = view
				return m, cmd
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
	} else {
		card = styles.PaddingV1.Render(card)
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

	containerWidth := min(width-6, 100)
	containerHeight := min(height, 25)

	container := styles.Merge([]lipgloss.Style{
		styles.BorderStyle,
		styles.Width(containerWidth),
		styles.Height(containerHeight),
	})

	innerWidth := containerWidth - 2
	innerHeight := containerHeight - 2

	var view string
	switch m.focus {
	case gameplay:
		view = m.GameplayView.View(innerWidth, innerHeight)
	case display:
		view = m.DisplayView.View(innerWidth, innerHeight)
	case audio:
		view = m.AudioView.View(innerWidth, innerHeight)
	case controls:
		view = m.ControlsView.View(innerWidth, innerHeight)
	}

	title := lipgloss.JoinVertical(lipgloss.Center,
		list,
		container.Render(view),
		footer,
	)

	return title
}
