package game

import (
	"math/rand"

	state "terminal.minesweeper/shared"
)

type Coords struct {
	X int
	Y int
}

type Tile struct {
	state    state.TileState
	adjacent int
}

type Board struct {
	tiles    [][]Tile
	time     int
	started  bool
	complete bool
	defeated bool

	width      int
	height     int
	mine_count int
}

func GenerateBoard(width int, height int, mine_count int) *Board {
	tiles := make([][]Tile, height)
	for i := range tiles {
		tiles[i] = make([]Tile, width)
	}

	return &Board{
		tiles:      tiles,
		width:      width,
		height:     height,
		mine_count: mine_count,
	}
}

func (b *Board) Populate(coord Coords) {
	// Preclear starting tiles
	start := b.GetTile(coord)
	start.state = state.TileOpen
	start.adjacent = 0

	for _, ncoords := range b.GetNeighbors(coord) {
		b.SetTileState(ncoords, state.TileOpen)
	}

	// Randomly place all mines
	local_count := 0
	for local_count < b.mine_count {
		x, y := rand.Intn(b.width), rand.Intn(int(b.height))

		if t := b.GetTile(Coords{X: x, Y: y}); t != nil &&
			t.state != state.TileOpen && t.state != state.MineClosed {
			t.state = state.MineClosed

			for _, ncoords := range b.GetNeighbors(Coords{X: x, Y: y}) {
				b.GetTile(ncoords).adjacent++
			}

			local_count++
		}
	}
}

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) GetTile(coord Coords) *Tile {
	if coord.X < 0 || coord.X >= b.width || coord.Y < 0 || coord.Y >= b.height {
		return nil
	}
	return &b.tiles[coord.Y][coord.X]
}

func (b *Board) GetNeighbors(coord Coords) []Coords {
	neighbors := []Coords{}
	for y := coord.Y - 1; y <= coord.Y+1; y++ {
		for x := coord.X - 1; x <= coord.X+1; x++ {
			if b.GetTile(Coords{X: x, Y: y}) != nil {
				neighbors = append(neighbors, Coords{X: x, Y: y})
			}
		}
	}

	return neighbors
}

func (b *Board) GetTileState(coord Coords) state.TileState {
	return b.GetTile(coord).state
}

func (b *Board) SetTileState(coord Coords, tileState state.TileState) {
	b.GetTile(coord).state = tileState
}

func (b *Board) Adjacent(coord Coords) int {
	return b.GetTile(coord).adjacent
}

func (b *Board) Flag(coord Coords) {
	switch b.GetTileState(coord) {
	case state.TileFlagged:
		b.SetTileState(coord, state.MineClosed)
	case state.MineClosed:
		b.SetTileState(coord, state.TileFlagged)
	case state.TileClosed:
		b.SetTileState(coord, state.TileFlaggedWrong)
	case state.TileFlaggedWrong:
		b.SetTileState(coord, state.TileClosed)
	}
}

func (b *Board) OpenTile(coord Coords) {
	if b.GetTileState(coord) == state.TileFlagged {
		return
	}

	if b.GetTileState(coord) == state.MineClosed {
		b.SetTileState(coord, state.MineHit)
		b.defeated = true
		b.Complete()
		return
	}

	if !b.started {
		b.Populate(coord)
		b.started = true

		for _, ncoords := range b.GetNeighbors(coord) {
			if b.Adjacent(ncoords) > 0 {
				continue
			}
			for _, acoords := range b.GetNeighbors(ncoords) {
				b.OpenSafeTile(acoords)
			}
		}
	} else {
		b.OpenSafeTile(coord)
	}

}

func (b *Board) OpenSafeTile(coord Coords) {
	if t := b.GetTileState(coord); t == state.TileOpen || t == state.TileFlagged {
		return
	}
	b.SetTileState(coord, state.TileOpen)
	if b.Adjacent(coord) > 0 {
		if b.IsBoardSolved() {
			b.Complete()
		}
		return
	}

	for _, ncoords := range b.GetNeighbors(coord) {
		b.OpenSafeTile(ncoords)
	}

	if b.IsBoardSolved() {
		b.Complete()
	}
}

func (b *Board) IsBoardSolved() bool {
	for _, row := range b.tiles {
		for _, tile := range row {
			if tile.state == state.TileClosed {
				return false
			}
		}
	}
	return true
}

func (b *Board) RevealBoard() {
	for y := range b.GetHeight() {
		for x := range b.GetWidth() {
			switch b.GetTileState(Coords{X: x, Y: y}) {
			case state.TileClosed:
				b.SetTileState(Coords{X: x, Y: y}, state.TileOpen)
			case state.MineClosed:
				b.SetTileState(Coords{X: x, Y: y}, state.MineRevealed)
			}
		}
	}
}

func (b *Board) Complete() {
	b.RevealBoard()
	b.complete = true
}

func (b *Board) IsComplete() bool {
	return b.complete
}
