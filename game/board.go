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
	mine     bool
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
		// for j := range tiles[i] {
		// 	tiles[i][j].adjacent = rand.Intn(9)
		// }
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

	neighbors := b.GetNeighbors(coord)

	for _, n := range neighbors {
		n.state = state.TileOpen
	}

	// Randomly place all mines
	local_count := 0
	for {
		if local_count == b.mine_count {
			break
		}
		x, y := rand.Intn(b.width), rand.Intn(int(b.height))

		if t := b.GetTile(Coords{X: x, Y: y}); t != nil &&
			t.state != state.TileOpen && !t.mine {
			t.state = state.TileClosed
			t.mine = true

			adjacents := b.GetNeighbors(coord)
			for _, n := range adjacents {
				n.adjacent++
			}

			local_count++
		}
	}

	// Game has been started
	b.started = true
}

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) GetTile(coord Coords) *Tile {
	if coord.X < 0 || coord.X >= b.width || coord.Y < 0 || coord.Y > b.height {
		return nil
	}
	return &b.tiles[coord.Y][coord.X]
}

func (b *Board) GetNeighbors(coord Coords) []*Tile {
	neighbors := []*Tile{}
	for y := coord.Y - 1; y <= coord.Y+1; y++ {
		for x := coord.X - 1; y <= coord.X+1; x++ {
			if tile := b.GetTile(Coords{X: x, Y: y}); tile != nil &&
				y != coord.Y && x != coord.X {
				neighbors = append(neighbors, tile)
			}
		}
	}

	return neighbors
}

func (b *Board) GetTileState(coord Coords) state.TileState {
	return b.GetTile(coord).state
}

func (b *Board) SetTileState(coord Coords, tileState state.TileState) {
	switch tileState {
	case state.TileOpen:
		b.OpenTile(coord)
	default:
		b.GetTile(coord).state = tileState
	}
}

func (b *Board) OpenTile(coord Coords) {
	if !b.started {
		b.Populate(coord)
	}

	// neighbors := b.GetNeighbors(coord)
	// for _, n := range neighbors {

	// }
}

func (b *Board) Adjacent(coord Coords) int {
	return b.GetTile(coord).adjacent
}

func (b *Board) IsComplete() bool {
	return b.complete
}
