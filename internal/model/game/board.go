package game

type CellState int

const (
	Empty CellState = iota
	Black
	White
)

type Board struct {
	ID    int           `json:"id" bson:"_id"`
	Size  int           `json:"size" bson:"size"`
	Cells [][]CellState `json:"cells" bson:"cells"`
}

func (b Board) GetSize() int {
	return b.Size
}

func NewBoard(size int) *Board {
	board := &Board{
		Size:  size,
		Cells: make([][]CellState, size),
	}

	for i := 0; i < size; i++ {
		board.Cells[i] = make([]CellState, size)
		for j := 0; j < size; j++ {
			board.Cells[i][j] = Empty
		}
	}

	return board
}
