package fields

import (
	"fmt"
	"reflect"
	"strconv"

	tea "charm.land/bubbletea/v2"
	"terminal.minesweeper/config"
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

func toString[T any](v T) string {
	val := reflect.ValueOf(v)
	switch val.Kind() {
	case reflect.Slice:
		res := make([]string, val.Len())
		for i := 0; i < val.Len(); i++ {
			e := val.Index(i)
			res[i] = e.String()
		}

		return config.FormatHelp(res)
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

func truncate[T any](v T) T {
	val := reflect.ValueOf(v)
	var result any
	switch val.Kind() {
	// Special case: for arrays the user might be trying to bind
	// the backspace key to a control
	case reflect.Slice:
		val = reflect.Append(val, reflect.ValueOf("backspace"))
		result = val.Interface()
	case reflect.Int:
		result = int(val.Int() / 10)
	}

	return result.(T)
}

func appendTo[T any](v T, key string) T {
	val := reflect.ValueOf(v)

	var result any
	switch val.Kind() {
	case reflect.Slice:
		val = reflect.Append(val, reflect.ValueOf(key))
		result = val.Interface()
	case reflect.Int:
		valString := toString(v) + key
		newVal, err := strconv.Atoi(valString)
		if err != nil {
			return v
		}
		result = newVal
	}

	return result.(T)
}
