package game

type Coords struct {
	X int
	Y int
}

type Tile struct {
	revealed bool
	flagged  bool
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
	b.GetTile(coord).adjacent = 0
	for y := coord.Y - 1; y <= coord.Y+1; y++ {
		for x := coord.X - 1; y <= coord.X+1; x++ {
			if tile := b.GetTile(Coords{X: x, Y: y}); tile != nil {
				tile.revealed = true
			}
		}
	}
	// Randomly place all mines
	local_count := 0
	for {
		if local_count == b.mine_count {
			break
		}
		// x, y := rand.Intn(b.width), rand.Intn(int(b.height))
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

func (b *Board) Flag(coord Coords) {
	tile := b.GetTile(coord)
	if !tile.revealed {
		tile.flagged = !tile.flagged
	}
}

func (b *Board) IsFlagged(coord Coords) bool {
	return b.GetTile(coord).flagged
}

func (b *Board) Reveal(coord Coords) {
	if !b.started {
		b.Populate(coord)
	}
}

func (b *Board) IsRevealed(coord Coords) bool {
	return b.GetTile(coord).revealed
}

func (b *Board) Adjacent(coord Coords) int {
	return b.GetTile(coord).adjacent
}

func (b *Board) IsMine(coord Coords) bool {
	return b.GetTile(coord).mine
}

func (b *Board) IsComplete() bool {
	return b.complete
}
