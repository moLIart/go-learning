package game

type CellState int

const (
	Empty CellState = iota
	Black
	White
)

type Board struct {
	size  int
	cells [][]CellState
}

func (b Board) Size() int {
	return b.size
}

func NewBoard(size int) *Board {
	// create
	board := &Board{
		size:  size,
		cells: make([][]CellState, size),
	}

	// initialize
	for i := 0; i < size; i++ {
		board.cells[i] = make([]CellState, size)

		for j := 0; j < size; j++ {
			board.cells[i][j] = Empty
		}
	}

	return board
}
