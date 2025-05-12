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

	playerIDCounter int
	roomIDCounter   int
	boardIDCounter  int
	gameIDCounter   int
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
		playerIDCounter++
		e.ID = playerIDCounter
		players = append(players, e)
		saveToFile(playersFile, players)
	case *room.Room:
		roomIDCounter++
		e.ID = roomIDCounter
		rooms = append(rooms, e)
		saveToFile(roomsFile, rooms)
	case *game.Board:
		boardIDCounter++
		e.ID = boardIDCounter
		boards = append(boards, e)
		saveToFile(boardsFile, boards)
	case *game.Game:
		gameIDCounter++
		e.ID = gameIDCounter
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

func GetPlayers() []*room.Player {
	mu.Lock()
	defer mu.Unlock()
	return players
}

func GetRooms() []*room.Room {
	mu.Lock()
	defer mu.Unlock()
	return rooms
}

func GetBoards() []*game.Board {
	mu.Lock()
	defer mu.Unlock()
	return boards
}

func GetGames() []*game.Game {
	mu.Lock()
	defer mu.Unlock()
	return games
}

func GetPlayerByID(id int) *room.Player {
	mu.Lock()
	defer mu.Unlock()
	for _, player := range players {
		if player.ID == id {
			return player
		}
	}
	return nil
}

func GetRoomByID(id int) *room.Room {
	mu.Lock()
	defer mu.Unlock()
	for _, room := range rooms {
		if room.ID == id {
			return room
		}
	}
	return nil
}

func GetBoardByID(id int) *game.Board {
	mu.Lock()
	defer mu.Unlock()
	for _, board := range boards {
		if board.ID == id {
			return board
		}
	}
	return nil
}

func GetGameByID(id int) *game.Game {
	mu.Lock()
	defer mu.Unlock()
	for _, game := range games {
		if game.ID == id {
			return game
		}
	}
	return nil
}

func DeletePlayerByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, player := range players {
		if player.ID == id {
			players = append(players[:i], players[i+1:]...)
			saveToFile(playersFile, players)
			return true
		}
	}
	return false
}

func DeleteRoomByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, room := range rooms {
		if room.ID == id {
			rooms = append(rooms[:i], rooms[i+1:]...)
			saveToFile(roomsFile, rooms)
			return true
		}
	}
	return false
}

func DeleteBoardByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, board := range boards {
		if board.ID == id {
			boards = append(boards[:i], boards[i+1:]...)
			saveToFile(boardsFile, boards)
			return true
		}
	}
	return false
}

func DeleteGameByID(id int) bool {
	mu.Lock()
	defer mu.Unlock()
	for i, game := range games {
		if game.ID == id {
			games = append(games[:i], games[i+1:]...)
			saveToFile(gamesFile, games)
			return true
		}
	}
	return false
}
