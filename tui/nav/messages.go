package nav

type RouteState int

const (
	Title RouteState = iota
	Sweeper
	Settings
)

type Navigate struct {
	Route   RouteState
	Payload struct{}
}
