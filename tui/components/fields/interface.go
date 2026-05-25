package fields

import (
	"fmt"
	"reflect"
	"strconv"

	tea "charm.land/bubbletea/v2"
)

type Field interface {
	Update(tea.Msg) (Field, tea.Cmd)
	View(width, height int, focused bool) string
}

func toString[T any](v T) string {

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
