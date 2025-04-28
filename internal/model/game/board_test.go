package game

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	size := 5
	board := NewBoard(size)

	if board.Size() != size {
		t.Errorf("expected board size %d, got %d", size, board.Size())
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board.cells[i][j] != Empty {
				t.Errorf("expected cell (%d, %d) to be Empty, got %d", i, j, board.cells[i][j])
			}
		}
	}
}
