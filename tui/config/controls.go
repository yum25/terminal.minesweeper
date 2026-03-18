package config

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

type UserControlsMap struct {
	Up     key.Binding
	Down   key.Binding
	Left   key.Binding
	Right  key.Binding
	Select key.Binding
	Quit   key.Binding
}

type GameControlsMap struct {
	UserControlsMap
	Flag    key.Binding
	Menu    key.Binding
	Restart key.Binding
}

var UserKeyMap = UserControlsMap{
	Up: key.NewBinding(
		key.WithKeys("w", "up"),
		key.WithHelp("↑/w", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("s", "down"),
		key.WithHelp("↓/s", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("a", "left"),
		key.WithHelp("←/a", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("d", "right"),
		key.WithHelp("→/d", "move right"),
	),
	Select: key.NewBinding(
		key.WithKeys("enter", "space"),
		key.WithHelp("enter/space", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
}

var GameKeyMap = GameControlsMap{
	UserControlsMap: UserControlsMap{
		Up: key.NewBinding(
			key.WithKeys("w", "up"),
			key.WithHelp("↑/w", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("s", "down"),
			key.WithHelp("↓/s", "move down"),
		),
		Left: key.NewBinding(
			key.WithKeys("a", "left"),
			key.WithHelp("←/a", "move left"),
		),
		Right: key.NewBinding(
			key.WithKeys("d", "right"),
			key.WithHelp("→/d", "move right"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter", "space"),
			key.WithHelp("enter/space", "select"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "quit"),
		),
	},
	Flag: key.NewBinding(
		key.WithKeys("f"),
		key.WithHelp("f", "flag"),
	),
	Menu: key.NewBinding(
		key.WithKeys("m"),
		key.WithHelp("m", "menu"),
	),
	Restart: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "restart"),
	),
}

func (u UserControlsMap) ShortHelp() []key.Binding {
	return []key.Binding{u.Up, u.Down, u.Select, u.Quit}
}

func (u UserControlsMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{u.Up, u.Down, u.Select, u.Quit}}
}

func (g GameControlsMap) ShortHelp() []key.Binding {
	return []key.Binding{g.Up, g.Down, g.Left, g.Right, g.Select, g.Flag, g.Menu}
}

func (g GameControlsMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{g.Up, g.Down, g.Left, g.Right}, {g.Select, g.Flag, g.Menu, g.Restart}}
}

type ControlMap interface {
	ShortHelp() []key.Binding
	FullHelp() [][]key.Binding
}

func RenderHelp(c ControlMap) string {
	return help.New().View(c)
}
