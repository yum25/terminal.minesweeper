package game

type Tile struct {
	revealed bool
	flagged  bool
	mine     bool
	adjacent int
}

type Board struct {
	tiles    [][]Tile
	time     int
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

func (b *Board) Populate() {

}

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) GetTile(x int, y int) *Tile {
	return &b.tiles[y][x]
}

func (b *Board) Flag(x int, y int) {
	tile := b.GetTile(y, x)
	if !tile.revealed {
		tile.flagged = !tile.flagged
	}
}

func (b *Board) IsFlagged(x int, y int) bool {
	return b.GetTile(y, x).flagged
}
