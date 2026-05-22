package animations

import (
	"math"
	"time"

	tea "charm.land/bubbletea/v2"
)

type tickMsg struct {
	generation int
}

type Animation struct {
	generation int
	elapsed    float64
	duration   float64
	easing     func(float64) float64
	active     bool
}

func MakeAnimation(duration float64, easing func(float64) float64) *Animation {
	return &Animation{
		duration: duration,
		easing:   easing,
	}
}

func (a *Animation) Start() tea.Cmd {
	a.generation++
	a.elapsed = 0
	a.active = true

	return a.tick()
}

func (a *Animation) Stop() {
	a.generation++
	a.active = false
}

func (a *Animation) Progress() float64 {
	if a.duration == 0 {
		return 1
	}

	return a.easing(math.Min(a.elapsed/a.duration, 1.0))
}

func (a *Animation) tick() tea.Cmd {
	return tea.Tick(time.Millisecond*16, func(t time.Time) tea.Msg {
		return tickMsg{generation: a.generation}
	})
}
