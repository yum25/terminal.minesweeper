package state

type TileState int

const (
	TileClosed TileState = iota
	TileOpen
	TileFlagged
	TileFlaggedWrong
	MineHit
	MineRevealed
)
