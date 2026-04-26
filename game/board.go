package game

import (
	"fmt"
	"log"
	"math/rand"

	state "terminal.minesweeper/shared"
	"terminal.minesweeper/tui/config"
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
	seconds  int
	started  bool
	complete bool
	defeated bool

	width      int
	height     int
	mine_count int
	flag_count int

	lives_count int
	lives_left  int
}

type TileOutOfBoundsError struct {
	coord Coords
}

type UnsupportedBoardDimensions struct {
	width  int
	height int
}

type ImpossibleMineCount struct {
	mine_count int
	width      int
	height     int
}

func (e *TileOutOfBoundsError) Error() string {
	return fmt.Sprintf("Attempted to access tile out of bounds at Coords{X: %d, Y: %d}",
		e.coord.X, e.coord.Y)
}

func (e *UnsupportedBoardDimensions) Error() string {
	return fmt.Sprintf("Unsupported board dimensions with width: %d, height: %d",
		e.width, e.height)
}

func (e *ImpossibleMineCount) Error() string {
	return fmt.Sprintf("Impossible mine count of %d mines with given dimensions %dx%d",
		e.mine_count, e.width, e.height)
}

func GenerateBoard(config *config.BoardConfig) *Board {
	tiles := make([][]Tile, config.Height)
	for i := range tiles {
		tiles[i] = make([]Tile, config.Width)
	}

	return &Board{
		tiles:      tiles,
		width:      config.Width,
		height:     config.Height,
		mine_count: config.MineCount,
	}
}

func (b *Board) Populate(coord Coords) {
	// Reset all tiles
	for y, row := range b.tiles {
		for x := range row {
			b.SetTileState(Coords{x, y}, state.TileClosed)
		}
	}

	// Preclear starting tiles
	b.SetTileState(coord, state.TileOpen)
	b.SetAdjacent(coord, 0)

	for _, ncoords := range b.GetNeighbors(coord) {
		b.SetTileState(ncoords, state.TileOpen)
	}

	// Randomly place all mines
	local_count := 0
	for local_count < b.mine_count {
		x, y := rand.Intn(b.width), rand.Intn(b.height)

		if s := b.GetTileState(Coords{X: x, Y: y}); s != state.TileOpen &&
			s != state.MineClosed {
			b.SetTileState(Coords{X: x, Y: y}, state.MineClosed)

			for _, ncoords := range b.GetNeighbors(Coords{X: x, Y: y}) {
				b.SetAdjacent(ncoords, b.GetAdjacent(ncoords)+1)
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

func (b *Board) GetFlagCount() int {
	return b.flag_count
}

func (b *Board) GetMineCount() int {
	return b.mine_count
}

func (b *Board) GetTile(coord Coords) (*Tile, *TileOutOfBoundsError) {
	if coord.X < 0 || coord.X >= b.width || coord.Y < 0 || coord.Y >= b.height {
		return nil, &TileOutOfBoundsError{}
	}
	return &b.tiles[coord.Y][coord.X], nil
}

func (b *Board) GetTileState(coord Coords) state.TileState {
	t, err := b.GetTile(coord)
	if err != nil {
		log.Fatal(err.Error())
	}
	return t.state
}

func (b *Board) SetTileState(coord Coords, tileState state.TileState) {
	t, err := b.GetTile(coord)
	if err != nil {
		log.Printf("WARNING: Board.SetTileState produced error %s", err.Error())
		return
	}

	t.state = tileState
}

func (b *Board) GetNeighbors(coord Coords) []Coords {
	neighbors := []Coords{}
	for y := coord.Y - 1; y <= coord.Y+1; y++ {
		for x := coord.X - 1; x <= coord.X+1; x++ {
			if _, err := b.GetTile(Coords{X: x, Y: y}); err == nil {
				neighbors = append(neighbors, Coords{X: x, Y: y})
			}
		}
	}

	return neighbors
}

func (b *Board) GetAdjacent(coord Coords) int {
	t, err := b.GetTile(coord)
	if err != nil {
		log.Printf("WARNING: Board.Adjacent produced error %s", err.Error())
		return 0
	}
	return t.adjacent
}

func (b *Board) SetAdjacent(coord Coords, count int) {
	t, err := b.GetTile(coord)
	if err != nil {
		log.Fatal(err.Error())
	}
	t.adjacent = count
}

func (b *Board) SetFlag(coord Coords) {
	switch b.GetTileState(coord) {
	case state.TileFlagged:
		b.SetTileState(coord, state.MineClosed)
		b.flag_count--

	case state.MineClosed:
		b.SetTileState(coord, state.TileFlagged)
		b.flag_count++

	case state.TileClosed:
		b.SetTileState(coord, state.TileFlaggedWrong)
		b.flag_count++

	case state.TileFlaggedWrong:
		b.SetTileState(coord, state.TileClosed)
		b.flag_count--
	}
}

func (b *Board) OpenTile(coord Coords) {
	if b.GetTileState(coord) == state.TileFlagged {
		return
	}

	if b.GetTileState(coord) == state.MineClosed {
		b.SetTileState(coord, state.MineHit)
		b.defeated = true
		return
	}

	if !b.started {
		b.Populate(coord)
		b.started = true

		for _, ncoords := range b.GetNeighbors(coord) {
			if b.GetAdjacent(ncoords) > 0 {
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
	if b.GetAdjacent(coord) > 0 {
		return
	}

	for _, ncoords := range b.GetNeighbors(coord) {
		b.OpenSafeTile(ncoords)
	}
}

func (b *Board) IsStarted() bool {
	return b.started
}

func (b *Board) GetTime() int {
	return b.seconds
}

func (b *Board) Tick() {
	b.seconds++
}

func (b *Board) IsBoardSolved() bool {
	for _, row := range b.tiles {
		for _, tile := range row {
			if tile.state == state.TileClosed || tile.state == state.TileFlaggedWrong {
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

func (b *Board) CheckIsComplete() {
	if b.IsBoardSolved() || b.defeated {
		b.RevealBoard()
		b.complete = true
	}
}

func (b *Board) IsComplete() bool {
	return b.complete
}
