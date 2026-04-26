package state

type TileState int

const (
	TileClosed TileState = iota
	TileOpen
	TileFlagged
	TileFlaggedWrong
	MineClosed
	MineHit

	// Post game tile states
	MineRevealed
	TileWin
	TileFlaggedWrongRevealed
	MineWin
)
