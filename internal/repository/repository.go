package repository

import (
	"sync"

	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/model/room"
)

var (
	players []room.Player
	rooms   []room.Room
	boards  []game.Board
	games   []game.Game
	mu      sync.Mutex
)

type Entity interface{}

func AddEntity(entity Entity) {
	mu.Lock()
	defer mu.Unlock()

	switch e := entity.(type) {
	case *room.Room:
		rooms = append(rooms, *e)
	case *room.Player:
		players = append(players, *e)
	case *game.Board:
		boards = append(boards, *e)
	case *game.Game:
		games = append(games, *e)
	}
}

func GetNumOfEntities() int {
	mu.Lock()
	defer mu.Unlock()

	return len(players) + len(rooms) + len(boards) + len(games)
}

func GetPlayersCount() int {
	mu.Lock()
	defer mu.Unlock()
	return len(players)
}

func GetRoomsCount() int {
	mu.Lock()
	defer mu.Unlock()
	return len(rooms)
}

func GetBoardsCount() int {
	mu.Lock()
	defer mu.Unlock()
	return len(boards)
}

func GetGamesCount() int {
	mu.Lock()
	defer mu.Unlock()
	return len(games)
}
