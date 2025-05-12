package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/moLIart/go-course/internal/handlers"
	"github.com/moLIart/go-course/internal/repository"

	"github.com/julienschmidt/httprouter"
)

func main() {
	repository.LoadData()

	router := httprouter.New()
	router.POST("/players", handlers.CreatePlayerHandler)
	router.GET("/players", handlers.GetPlayersHandler)
	router.GET("/players/:id", handlers.GetPlayerByIDHandler)
	router.PUT("/players/:id", handlers.UpdatePlayerHandler)
	router.DELETE("/players/:id", handlers.DeletePlayerByIDHandler)

	router.POST("/rooms", handlers.CreateRoomHandler)
	router.GET("/rooms", handlers.GetRoomsHandler)
	router.GET("/rooms/:id", handlers.GetRoomByIDHandler)
	router.PUT("/rooms/:id", handlers.UpdateRoomHandler)
	router.DELETE("/rooms/:id", handlers.DeleteRoomHandler)

	router.POST("/boards", handlers.CreateBoardHandler)
	router.GET("/boards", handlers.GetBoardsHandler)
	router.GET("/boards/:id", handlers.GetBoardByIDHandler)
	router.PUT("/boards/:id", handlers.UpdateBoardHandler)
	router.DELETE("/boards/:id", handlers.DeleteBoardHandler)

	router.POST("/games", handlers.CreateGameHandler)
	router.GET("/games", handlers.GetGamesHandler)
	router.GET("/games/:id", handlers.GetGameByIDHandler)
	router.DELETE("/games/:id", handlers.DeleteGameHandler)

	srv := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-signalChan
		log.Println("Shutting down server...")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Fatalf("Server forced to shutdown: %s", err)
		}
	}()

	log.Println("Starting server on :8081")
	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatalf("Could not start server: %s", err)
	}

	wg.Wait()
	log.Println("Server gracefully stopped")
}
