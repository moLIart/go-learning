package repository

import (
	"encoding/json"
	"os"
	"sync"

	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/model/room"
)

var (
	players []*room.Player
	rooms   []*room.Room
	boards  []*game.Board
	games   []*game.Game
	mu      sync.Mutex
)

const (
	playersFile = "players.json"
	roomsFile   = "rooms.json"
	boardsFile  = "boards.json"
	gamesFile   = "games.json"
)

type Entity interface{}

func AddEntity(entity Entity) {
	mu.Lock()
	defer mu.Unlock()

	switch e := entity.(type) {
	case *room.Player:
		players = append(players, e)
		saveToFile(playersFile, players)
	case *room.Room:
		rooms = append(rooms, e)
		saveToFile(roomsFile, rooms)
	case *game.Board:
		boards = append(boards, e)
		saveToFile(boardsFile, boards)
	case *game.Game:
		games = append(games, e)
		saveToFile(gamesFile, games)
	}
}

func saveToFile(filename string, data interface{}) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(data); err != nil {
		panic(err)
	}
}

func LoadData() {
	loadFromFile(playersFile, &players)
	loadFromFile(roomsFile, &rooms)
	loadFromFile(boardsFile, &boards)
	loadFromFile(gamesFile, &games)
}

func loadFromFile(filename string, target interface{}) {
	file, err := os.Open(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(target); err != nil {
		panic(err)
	}
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
