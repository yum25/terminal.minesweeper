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
