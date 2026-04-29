package config

import (
	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

type ControlPreset string

const (
	DefaultControls ControlPreset = "DEFAULT"
	VimControls     ControlPreset = "VIM"
	Custom          ControlPreset = "CUSTOM"
)

type UserControlsMap struct {
	Up     key.Binding `json:"Up"`
	Down   key.Binding `json:"Down"`
	Left   key.Binding `json:"Left"`
	Right  key.Binding `json:"Right"`
	Select key.Binding `json:"Select"`
	Quit   key.Binding `json:"Quit"`
}

type GameControlsMap struct {
	UserControlsMap
	Flag    key.Binding `json:"Flag"`
	Menu    key.Binding `json:"Menu"`
	Restart key.Binding `json:"Restart"`
}

func (u UserControlsMap) ShortHelp() []key.Binding {
	return []key.Binding{u.Up, u.Down, u.Select, u.Quit}
}

func (u UserControlsMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{u.Up, u.Down, u.Select, u.Quit}}
}

func (g GameControlsMap) ShortHelp() []key.Binding {
	return []key.Binding{g.Select, g.Flag, g.Menu, g.Restart}
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

var DEFAULT_USERKEYMAP = UserControlsMap{
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

var DEFAULT_GAMEKEYMAP = GameControlsMap{
	UserControlsMap: DEFAULT_USERKEYMAP,
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

var VIM_USERKEYMAP = UserControlsMap{
	Up: key.NewBinding(
		key.WithKeys("k", "up"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("j", "down"),
		key.WithHelp("↓/j", "move down"),
	),
	Left: key.NewBinding(
		key.WithKeys("h", "left"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("l", "right"),
		key.WithHelp("→/l", "move right"),
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

var VIM_GAMEKEYMAP = GameControlsMap{
	UserControlsMap: VIM_USERKEYMAP,
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

var UserKeyMap = DEFAULT_USERKEYMAP
var GameKeyMap = DEFAULT_GAMEKEYMAP
