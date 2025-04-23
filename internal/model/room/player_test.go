package room

import (
	"testing"
)

func TestNewPlayer(t *testing.T) {
	name := "Alice"
	player := NewPlayer(name)

	if player.GetName() != name {
		t.Errorf("expected player name to be %s, got %s", name, player.GetName())
	}
}

func TestGetName(t *testing.T) {
	player := NewPlayer("Bob")

	if player.GetName() != "Bob" {
		t.Errorf("expected player name to be Bob, got %s", player.GetName())
	}
}

func TestSetName(t *testing.T) {
	player := NewPlayer("Charlie")
	newName := "Dave"

	player.SetName(newName)

	if player.GetName() != newName {
		t.Errorf("expected player name to be %s, got %s", newName, player.GetName())
	}
}
