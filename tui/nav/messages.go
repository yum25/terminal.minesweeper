package nav

type RouteState int

const (
	Title RouteState = iota
	Sweeper
)

type Navigate struct {
	Route   RouteState
	Payload struct{}
}
