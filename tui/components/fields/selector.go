package fields

import (
	"slices"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/styles"
)

type SelectorModel[T comparable] struct {
	name     string
	value    *T
	prev     T
	options  []T
	cursor   int
	controls *config.UserControlsMap
}

func MakeSelectorModel[T comparable](
	name string,
	value *T,
	options []T,
	controls *config.UserControlsMap,
) SelectorModel[T] {
	return SelectorModel[T]{
		name:     name,
		value:    value,
		prev:     *value,
		options:  options,
		cursor:   slices.Index(options, *value),
		controls: controls,
	}
}

func (m SelectorModel[T]) GetName() string {
	return m.name
}

func (m SelectorModel[T]) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m SelectorModel[T]) Update(msg tea.Msg) (Field, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch {
		case key.Matches(msg, m.controls.Left):
			if m.cursor > 0 {
				m.cursor--
			} else {
				m.cursor = len(m.options) - 1
			}
			*m.value = m.options[m.cursor]
		case key.Matches(msg, m.controls.Right):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
			*m.value = m.options[m.cursor]
		case key.Matches(msg, m.controls.Select):
			m.prev = *m.value
		case key.Matches(msg, m.controls.Cancel):
			*m.value = m.prev
		}
	}

	return m, nil
}

func (m SelectorModel[T]) View(width, height int, state State) string {
	valWidth := 0
	for _, option := range m.options {
		localWidth := lipgloss.Width(toString(option)) + 2
		if localWidth > valWidth {
			valWidth = localWidth
		}
	}
	var style lipgloss.Style
	switch state {
	case Unfocused:
		style = styles.OptionStyle
	case Hover:
		style = styles.HoveredOptionStyle
	case Focused:
		style = styles.SelectedOptionStyle
	}
	option := styles.Merge([]lipgloss.Style{
		style,
		styles.Width(valWidth)},
	).Render(toString(*m.value))

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render(
		lipgloss.JoinHorizontal(lipgloss.Center,
			constants.ArrowLeftSymbol,
			option,
			constants.ArrowRightSymbol),
	)
}
