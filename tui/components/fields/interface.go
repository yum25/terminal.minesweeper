package fields

import (
	"fmt"
	"reflect"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

type State int

const (
	Unfocused State = iota
	Hover
	Focused
)

type Field interface {
	GetName() string
	Init() tea.Cmd
	Update(tea.Msg) (Field, tea.Cmd)
	View(width, height int, state State) string
}

func toString[T comparable](v T) string {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.String:
		return val.String()
	case reflect.Int:
		return strconv.FormatInt(val.Int(), 10)
	case reflect.Bool:
		return strconv.FormatBool(val.Bool())
	default:
		return fmt.Sprintf("%v", v)
	}
}
