package room

import (
	"testing"

	"github.com/moLIart/go-course/internal/model/game"
)

func TestNewRoom(t *testing.T) {
	code := "room123"
	room := NewRoom(code)

	if room.GetCode() != code {
		t.Errorf("expected room code to be %s, got %s", code, room.GetCode())
	}

	if room.GetGame() == nil {
		t.Error("expected game to be initialized, got nil")
	}

	if !room.IsEmpty() {
		t.Error("expected room to be empty initially")
	}
}

func TestAddPlayer(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")
	player3 := NewPlayer("Charlie")

	if !room.AddPlayer(player1) {
		t.Error("expected to successfully add player1")
	}

	if !room.AddPlayer(player2) {
		t.Error("expected to successfully add player2")
	}

	if room.AddPlayer(player3) {
		t.Error("expected to fail adding player3 as the room is full")
	}

	if !room.IsFull() {
		t.Error("expected room to be full after adding two players")
	}
}

func TestRemovePlayer(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")

	room.AddPlayer(player1)
	room.AddPlayer(player2)

	if !room.RemovePlayer(player1) {
		t.Error("expected to successfully remove player1")
	}

	if room.IsFull() {
		t.Error("expected room to not be full after removing a player")
	}

	if !room.RemovePlayer(player2) {
		t.Error("expected to successfully remove player2")
	}

	if !room.IsEmpty() {
		t.Error("expected room to be empty after removing all players")
	}
}

func TestGetCurrentPlayer(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")

	room.AddPlayer(player1)
	room.AddPlayer(player2)

	if room.GetCurrentPlayer() != player1 {
		t.Error("expected current player to be player1 (Black)")
	}

	room.GetGame().SwitchTurn()

	if room.GetCurrentPlayer() != player2 {
		t.Error("expected current player to be player2 (White)")
	}
}

func TestGetOpponent(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")

	room.AddPlayer(player1)
	room.AddPlayer(player2)

	if room.GetOpponent(player1) != player2 {
		t.Error("expected opponent of player1 to be player2")
	}

	if room.GetOpponent(player2) != player1 {
		t.Error("expected opponent of player2 to be player1")
	}
}

func TestGetPlayerByName(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")

	room.AddPlayer(player1)
	room.AddPlayer(player2)

	if room.GetPlayerByName("Alice") != player1 {
		t.Error("expected to find player1 by name Alice")
	}

	if room.GetPlayerByName("Bob") != player2 {
		t.Error("expected to find player2 by name Bob")
	}

	if room.GetPlayerByName("Charlie") != nil {
		t.Error("expected to not find a player with name Charlie")
	}
}

func TestGetPlayerByColor(t *testing.T) {
	room := NewRoom("room123")
	player1 := NewPlayer("Alice")
	player2 := NewPlayer("Bob")

	room.AddPlayer(player1)
	room.AddPlayer(player2)

	if room.GetPlayerByColor(game.Black) != player1 {
		t.Error("expected player1 to be associated with Black color")
	}

	if room.GetPlayerByColor(game.White) != player2 {
		t.Error("expected player2 to be associated with White color")
	}
}
