package room

import (
	"github.com/moLIart/go-course/internal/model/game"
)

type Room struct {
	ID      int        `json:"id" bson:"_id"`
	Code    string     `json:"code" bson:"code"`
	Players [2]*Player `json:"players" bson:"players"`
	Game    *game.Game `json:"game" bson:"game"`
}

func NewRoom(code string) *Room {
	return &Room{
		Code:    code,
		Players: [2]*Player{},
		Game:    nil,
	}
}

func (r *Room) GetCode() string {
	return r.Code
}

func (r *Room) GetGame() *game.Game {
	return r.Game
}

func (r *Room) SetGame(game *game.Game) {
	r.Game = game
}

func (r *Room) AddPlayer(player *Player) bool {
	if r.Players[0] == nil {
		r.Players[0] = player
		return true
	} else if r.Players[1] == nil {
		r.Players[1] = player
		return true
	}
	return false
}

func (r *Room) RemovePlayer(player *Player) bool {
	if r.Players[0] == player {
		r.Players[0] = nil
		return true
	} else if r.Players[1] == player {
		r.Players[1] = nil
		return true
	}
	return false
}

func (r *Room) IsFull() bool {
	return r.Players[0] != nil && r.Players[1] != nil
}

func (r *Room) IsEmpty() bool {
	return r.Players[0] == nil && r.Players[1] == nil
}

func (r *Room) GetCurrentPlayer() *Player {
	if r.Game.GetCurrentTurn() == game.Black {
		return r.Players[0]
	}
	return r.Players[1]
}

func (r *Room) GetOpponent(player *Player) *Player {
	if r.Players[0] == player {
		return r.Players[1]
	}
	return r.Players[0]
}

func (r *Room) GetPlayerByName(name string) *Player {
	for _, player := range r.Players {
		if player != nil && player.GetName() == name {
			return player
		}
	}
	return nil
}

func (r *Room) GetPlayerByColor(color game.CellState) *Player {
	if color == game.Black {
		return r.Players[0]
	}
	return r.Players[1]
}
