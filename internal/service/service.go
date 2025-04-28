package service

import (
	"fmt"
	"time"

	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/model/room"
	"github.com/moLIart/go-course/internal/repository"
)

func StartProcessing(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for range ticker.C {
		newRoom := room.NewRoom("room123")
		newGame := game.NewGame()
		newBoard := game.NewBoard(17)
		newPlayer := room.NewPlayer("player123")

		repository.AddEntity(newRoom)
		repository.AddEntity(newGame)
		repository.AddEntity(newBoard)
		repository.AddEntity(newPlayer)

		fmt.Println("Entities added to repository. Current size: ", repository.GetNumOfEntities())
	}
}
