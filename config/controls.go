package config

import (
	"strings"

	"charm.land/bubbles/v2/help"
	"charm.land/bubbles/v2/key"
)

type ControlPreset string

const (
	DefaultControls ControlPreset = "DEFAULT"
	VimControls     ControlPreset = "VIM"
	Custom          ControlPreset = "CUSTOM"
)

type Controls struct {
	Up      []string `json:"Up"`
	Down    []string `json:"Down"`
	Left    []string `json:"Left"`
	Right   []string `json:"Right"`
	Select  []string `json:"Select"`
	Menu    []string `json:"Menu"`
	Cancel  []string `json:"Cancel"`
	Quit    []string `json:"Quit"`
	Flag    []string `json:"Flag"`
	Restart []string `json:"Restart"`
}

type UserControlsMap struct {
	Up     key.Binding `json:"Up"`
	Down   key.Binding `json:"Down"`
	Left   key.Binding `json:"Left"`
	Right  key.Binding `json:"Right"`
	Select key.Binding `json:"Select"`
	Menu   key.Binding `json:"Menu"`
	Cancel key.Binding `json:"Cancel"`
	Quit   key.Binding `json:"Quit"`
}

type GameControlsMap struct {
	UserControlsMap
	Flag    key.Binding `json:"Flag"`
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

func FormatHelp(keys []string) string {
	return strings.Join(keys, "/")
}

func ToKeyBinding(keys []string, description string) key.Binding {
	return key.NewBinding(
		key.WithKeys(keys...),
		key.WithHelp(FormatHelp(keys), description),
	)
}

func (c *Controls) ToKeyMap() (UserControlsMap, GameControlsMap) {
	UserControls := UserControlsMap{
		Up:     ToKeyBinding(c.Up, "move up"),
		Down:   ToKeyBinding(c.Down, "move down"),
		Left:   ToKeyBinding(c.Left, "move left"),
		Right:  ToKeyBinding(c.Right, "move right"),
		Select: ToKeyBinding(c.Select, "select"),
		Menu:   ToKeyBinding(c.Menu, "menu"),
		Cancel: ToKeyBinding(c.Cancel, "cancel"),
		Quit:   ToKeyBinding(c.Quit, "quit"),
	}
	return UserControls, GameControlsMap{
		UserControlsMap: UserControls,
		Flag:            ToKeyBinding(c.Flag, "flag"),
		Restart:         ToKeyBinding(c.Restart, "restart"),
	}
}

func (c Config) FromKeyMap() Controls {
	return Controls{
		Up:      c.UserControls.Up.Keys(),
		Down:    c.UserControls.Down.Keys(),
		Left:    c.UserControls.Left.Keys(),
		Right:   c.UserControls.Right.Keys(),
		Select:  c.UserControls.Select.Keys(),
		Cancel:  c.UserControls.Cancel.Keys(),
		Quit:    c.UserControls.Quit.Keys(),
		Flag:    c.GameControls.Flag.Keys(),
		Menu:    c.GameControls.Menu.Keys(),
		Restart: c.GameControls.Restart.Keys(),
	}
}

var DEFAULT_CONTROLS = Controls{
	Up:      []string{"w", "up"},
	Down:    []string{"s", "down"},
	Left:    []string{"a", "left"},
	Right:   []string{"d", "right"},
	Select:  []string{"enter", "space"},
	Cancel:  []string{"esc"},
	Quit:    []string{"q", "ctrl+c"},
	Flag:    []string{"f"},
	Menu:    []string{"m"},
	Restart: []string{"r"},
}

var VIM_CONTROLS = Controls{
	Up:      []string{"k", "up"},
	Down:    []string{"j", "down"},
	Left:    []string{"h", "left"},
	Right:   []string{"l", "right"},
	Select:  []string{"enter", "space"},
	Cancel:  []string{"esc"},
	Quit:    []string{"q", "esc", "ctrl+c"},
	Flag:    []string{"f"},
	Menu:    []string{"m"},
	Restart: []string{"r"},
}
