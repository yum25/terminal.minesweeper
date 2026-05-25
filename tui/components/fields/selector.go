package fields

import (
	"log"
	"slices"

	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/constants"
	"terminal.minesweeper/tui/styles"
)

type SelectorModel[T comparable] struct {
	value    *T
	options  []T
	cursor   int
	controls *config.UserControlsMap
}

func MakeSelectorModel[T comparable](
	value *T,
	options []T,
	controls *config.UserControlsMap,
) SelectorModel[T] {
	return SelectorModel[T]{
		value:    value,
		options:  options,
		cursor:   slices.Index(options, *value),
		controls: controls,
	}
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
				log.Print(m.cursor)
				m.cursor--
				log.Print(m.cursor)
			} else {
				m.cursor = len(m.options) - 1
			}
		case key.Matches(msg, m.controls.Right):
			if m.cursor < len(m.options)-1 {
				m.cursor++
			} else {
				m.cursor = 0
			}
		case key.Matches(msg, m.controls.Select):
			*m.value = m.options[m.cursor]
		}
	}

	return m, nil
}

func (m SelectorModel[T]) View(width, height int, state State) string {
	var option string
	switch state {
	case Unfocused:
		option = styles.OptionStyle.Render(toString(*m.value))
	case Hover:
		option = styles.HoveredOptionStyle.Render(toString(*m.value))
	case Focused:
		option = styles.SelectedOptionStyle.Render(toString(m.options[m.cursor]))
	}

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
