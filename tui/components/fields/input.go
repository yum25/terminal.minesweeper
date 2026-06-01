package fields

import (
	"charm.land/bubbles/v2/key"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
	"terminal.minesweeper/config"
	"terminal.minesweeper/tui/styles"
)

type InputModel[T any] struct {
	name     string
	buffer   T
	value    *T
	label    string
	controls *config.UserControlsMap
}

func MakeInputModel[T any](
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
		case msg.String() == "backspace":
			m.buffer = truncate(m.buffer)
		case key.Matches(msg, m.controls.Select):
			*m.value = m.buffer
			var clear T
			m.buffer = clear
		case key.Matches(msg, m.controls.Cancel):
			var clear T
			m.buffer = clear
		default:
			m.buffer = appendTo(m.buffer, msg.String())
		}
	}

	return m, nil
}

func (m InputModel[T]) View(width, height int, state State) string {
	// TODO: Refactor UI logic into a single card interface that uses same
	// logic as subview setting cards
	var style lipgloss.Style
	var text T

	switch state {
	case Unfocused:
		style = styles.OptionStyle
		text = *m.value
	case Hover:
		style = styles.HoveredOptionStyle
		text = *m.value
	case Focused:
		style = styles.SelectedOptionStyle
		text = m.buffer
	}

	val := styles.IndentStyle.Render(toString(text))
	input := styles.Merge([]lipgloss.Style{
		style,
		styles.Width(
			lipgloss.Width(m.label) + lipgloss.Width(val) + 3),
	}).Render(lipgloss.JoinHorizontal(lipgloss.Center,
		m.label,
		" ",
		val,
	))

	switch state {
	case Unfocused:
		input = styles.PaddingV1.Render(input)
	case Hover:
		input = styles.AddHalfPixelBorder(input,
			styles.Merge([]lipgloss.Style{
				styles.Text(styles.White),
				styles.Width(lipgloss.Width(input)),
			}),
		)
	case Focused:
		input = styles.AddHalfPixelBorder(input,
			styles.Merge([]lipgloss.Style{
				styles.Text(styles.CursorColor),
				styles.Width(lipgloss.Width(input)),
			}),
		)
	}

	return styles.Merge([]lipgloss.Style{
		styles.AlignCenter,
		styles.PaddingH1,
	}).Render(input)
}
