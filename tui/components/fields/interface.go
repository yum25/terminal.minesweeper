package fields

import tea "charm.land/bubbletea/v2"

type Field interface {
	Update(tea.Msg) (Field, tea.Cmd)
}
