package views

type RouteState int

const (
	Title RouteState = iota
	Sweeper
)

type Navigate struct {
	route   RouteState
	payload struct{}
}
