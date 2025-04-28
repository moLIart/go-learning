package room

import (
	"github.com/moLIart/go-course/internal/model/game"
)

type Room struct {
	code    string
	players [2]*Player
	game    *game.Game
}

func NewRoom(code string) *Room {
	return &Room{
		code:    code,
		players: [2]*Player{},
		game:    game.NewGame(),
	}
}

func (r *Room) GetCode() string {
	return r.code
}

func (r *Room) GetGame() *game.Game {
	return r.game
}

func (r *Room) AddPlayer(player *Player) bool {
	if r.players[0] == nil {
		r.players[0] = player
		return true
	} else if r.players[1] == nil {
		r.players[1] = player
		return true
	}
	return false
}

func (r *Room) RemovePlayer(player *Player) bool {
	if r.players[0] == player {
		r.players[0] = nil
		return true
	} else if r.players[1] == player {
		r.players[1] = nil
		return true
	}
	return false
}

func (r *Room) IsFull() bool {
	return r.players[0] != nil && r.players[1] != nil
}

func (r *Room) IsEmpty() bool {
	return r.players[0] == nil && r.players[1] == nil
}

func (r *Room) GetCurrentPlayer() *Player {
	if r.game.GetCurrentTurn() == game.Black {
		return r.players[0]
	}
	return r.players[1]
}

func (r *Room) GetOpponent(player *Player) *Player {
	if r.players[0] == player {
		return r.players[1]
	}
	return r.players[0]
}

func (r *Room) GetPlayerByName(name string) *Player {
	for _, player := range r.players {
		if player != nil && player.GetName() == name {
			return player
		}
	}
	return nil
}

func (r *Room) GetPlayerByColor(color game.CellState) *Player {
	if color == game.Black {
		return r.players[0]
	}
	return r.players[1]
}
