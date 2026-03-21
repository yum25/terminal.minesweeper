package nav

type RouteState int
type Payload string

const (
	Title RouteState = iota
	Sweeper
	Settings
)

const (
	Play     Payload = "Play"
	Continue Payload = "Continue"

	New    Payload = "New"
	Paused Payload = "Paused"
)

type Navigate struct {
	Route   RouteState
	Payload Payload
}
