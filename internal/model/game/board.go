package game

type CellState int

const (
	Empty CellState = iota
	Black
	White
)

type Board struct {
	Size  int
	Cells [][]CellState
}

func (b Board) GetSize() int {
	return b.Size
}

func NewBoard(size int) *Board {
	// create
	board := &Board{
		Size:  size,
		Cells: make([][]CellState, size),
	}

	// initialize
	for i := 0; i < size; i++ {
		board.Cells[i] = make([]CellState, size)

		for j := 0; j < size; j++ {
			board.Cells[i][j] = Empty
		}
	}

	return board
}
