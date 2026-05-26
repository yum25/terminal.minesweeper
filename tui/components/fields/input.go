package fields

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/styles"
)

type InputModel[T comparable] struct {
	name     string
	value    *T
	buffer   T
	label    string
	controls *config.UserControlsMap
}

func MakeInputModel[T comparable](
	name string,
	value *T,
	label string,
	controls *config.UserControlsMap) InputModel[T] {
	return InputModel[T]{
		name:     name,
		value:    value,
		label:    label,
		controls: controls,
	}
}

func (m InputModel[T]) GetName() string {
	return m.name
}

func (m InputModel[T]) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m InputModel[T]) Update(msg tea.Msg) (Field, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Select):
			*m.value = m.buffer
		default:
		}
	}

	return m, nil
}

func (m InputModel[T]) View(width, height int, state State) string {
	var style lipgloss.Style
	var val string

	switch state {
	case Unfocused:
		style = styles.OptionStyle
		val = toString(*m.value)
	case Hover:
		style = styles.HoveredOptionStyle
		val = toString(*m.value)
	case Focused:
		style = styles.SelectedOptionStyle
		val = toString(m.buffer)
	}

	input := lipgloss.JoinHorizontal(lipgloss.Center,
		style.Render(m.label),
		styles.IndentStyle.Render(val),
	)

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render(input)
}
