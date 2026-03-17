package state

type TileState int

const (
	TileClosed TileState = iota
	TileOpen
	TileFlagged
	TileFlaggedWrong
	MineClosed
	MineHit
	MineRevealed

	// Post game tile states
	TileWin
	TileFlaggedWrongRevealed
	MineWin
)
