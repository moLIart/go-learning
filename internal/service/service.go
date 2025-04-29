package service

import (
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/model/room"
	"github.com/moLIart/go-course/internal/repository"
)

func StartProcessing(interval time.Duration) {
	dataChannel := make(chan interface{})
	var wg sync.WaitGroup

	// Горутина для генерации данных
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			var entity interface{}
			switch rand.Intn(4) {
			case 0:
				entity = room.NewRoom("room123")
			case 1:
				entity = game.NewGame()
			case 2:
				entity = game.NewBoard(17)
			case 3:
				entity = room.NewPlayer("player123")
			}

			dataChannel <- entity
		}
		close(dataChannel)
	}()

	// Горутина для добавления данных в репозиторий
	wg.Add(1)
	go func() {
		defer wg.Done()
		for entity := range dataChannel {
			repository.AddEntity(entity)
		}
	}()

	// Горутина для логирования добавленных структур
	wg.Add(1)
	go func() {
		defer wg.Done()
		ticker := time.NewTicker(200 * time.Millisecond)
		defer ticker.Stop()

		var prevCounts = map[string]int{
			"players": 0,
			"rooms":   0,
			"boards":  0,
			"games":   0,
		}

		for range ticker.C {
			currentCounts := map[string]int{
				"players": repository.GetPlayersCount(),
				"rooms":   repository.GetRoomsCount(),
				"boards":  repository.GetBoardsCount(),
				"games":   repository.GetGamesCount(),
			}

			for key, prevCount := range prevCounts {
				if currentCounts[key] > prevCount {
					log.Printf("New %s added: %d\n", key, currentCounts[key]-prevCount)
				}
			}

			prevCounts = currentCounts
		}
	}()

	wg.Wait()
}
