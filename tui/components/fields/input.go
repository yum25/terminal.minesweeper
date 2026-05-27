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
	prev     T
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
		prev:     *value,
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
		case msg.String() == "backspace":
			*m.value = truncate(*m.value)
		case key.Matches(msg, m.controls.Select):
			m.prev = *m.value
		case key.Matches(msg, m.controls.Cancel):
			*m.value = m.prev
		default:
			*m.value = appendTo(*m.value, msg.String())
		}
	}

	return m, nil
}

func (m InputModel[T]) View(width, height int, state State) string {
	// TODO: Refactor UI logic into a single card interface that uses same
	// logic as subview setting cards
	var style lipgloss.Style

	switch state {
	case Unfocused:
		style = styles.OptionStyle
	case Hover:
		style = styles.HoveredOptionStyle
	case Focused:
		style = styles.SelectedOptionStyle
	}

	val := styles.IndentStyle.Render(toString(*m.value))
	input := styles.Merge([]lipgloss.Style{
		style,
		styles.Width(
			lipgloss.Width(m.label) + lipgloss.Width(val) + 3),
	}).Render(lipgloss.JoinHorizontal(lipgloss.Center,
		m.label,
		" ",
		val,
	))

	if state == Hover {
		input = styles.AddHalfPixelBorder(input,
			styles.Merge([]lipgloss.Style{
				styles.Text(styles.White),
				styles.Width(lipgloss.Width(input)),
			}),
		)
	}

	if state == Focused {
		input = styles.AddHalfPixelBorder(input,
			styles.Merge([]lipgloss.Style{
				styles.Text(styles.CursorColor),
				styles.Width(lipgloss.Width(input)),
			}),
		)
	}

	return styles.Merge([]lipgloss.Style{
		styles.Width(width),
		styles.Height(height),
		styles.AlignCenter,
	}).Render(input)
}
