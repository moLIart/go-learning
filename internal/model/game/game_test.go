package game

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := NewGame()

	if game.GetCurrentTurn() != Black {
		t.Errorf("expected initial turn to be Black, got %d", game.GetCurrentTurn())
	}

	if game.GetStatus() != NotDecidedYet {
		t.Errorf("expected initial status to be NotDecidedYet, got %d", game.GetStatus())
	}

	if game.GetBoard() == nil {
		t.Error("expected board to be initialized, got nil")
	}
}

func TestNewGameWithSize(t *testing.T) {
	size := 7
	game := NewGame(WithSize(size))

	if game.GetBoard().GetSize() != size {
		t.Errorf("expected board size %d, got %d", size, game.GetBoard().GetSize())
	}
}

func TestSwitchTurn(t *testing.T) {
	game := NewGame()

	game.SwitchTurn()
	if game.GetCurrentTurn() != White {
		t.Errorf("expected turn to switch to White, got %d", game.GetCurrentTurn())
	}

	game.SwitchTurn()
	if game.GetCurrentTurn() != Black {
		t.Errorf("expected turn to switch back to Black, got %d", game.GetCurrentTurn())
	}
}

func TestSetAndGetStatus(t *testing.T) {
	game := NewGame()

	game.SetStatus(BlackWon)
	if game.GetStatus() != BlackWon {
		t.Errorf("expected status to be BlackWon, got %d", game.GetStatus())
	}

	game.SetStatus(Draw)
	if game.GetStatus() != Draw {
		t.Errorf("expected status to be Draw, got %d", game.GetStatus())
	}
}

func TestIsOver(t *testing.T) {
	game := NewGame()

	if game.IsOver() {
		t.Error("expected game to not be over initially")
	}

	game.SetStatus(BlackWon)
	if !game.IsOver() {
		t.Error("expected game to be over after setting status to BlackWon")
	}
}

func TestIsCurrentTurn(t *testing.T) {
	game := NewGame()

	if !game.IsCurrentTurnBlack() {
		t.Error("expected initial turn to be Black")
	}

	game.SwitchTurn()
	if !game.IsCurrentTurnWhite() {
		t.Error("expected turn to be White after switching")
	}
}
