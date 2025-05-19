package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/moLIart/go-course/internal/handlers"
	"github.com/moLIart/go-course/internal/middlewares"
	"github.com/moLIart/go-course/internal/repository"

	"github.com/julienschmidt/httprouter"
)

// @title						Test API Server
// @version					1
// @description			API Server
//
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	repository.LoadData()

	middlewares.SetJwtSecret(os.Getenv("JWT_SECRET"))

	router := httprouter.New()
	router.POST("/players", middlewares.JWTAuth(handlers.CreatePlayerHandler))
	router.GET("/players", handlers.GetPlayersHandler)
	router.GET("/players/:id", handlers.GetPlayerByIDHandler)
	router.PUT("/players/:id", middlewares.JWTAuth(handlers.UpdatePlayerHandler))
	router.DELETE("/players/:id", middlewares.JWTAuth(handlers.DeletePlayerByIDHandler))

	router.POST("/rooms", middlewares.JWTAuth(handlers.CreateRoomHandler))
	router.GET("/rooms", handlers.GetRoomsHandler)
	router.GET("/rooms/:id", handlers.GetRoomByIDHandler)
	router.PUT("/rooms/:id", middlewares.JWTAuth(handlers.UpdateRoomHandler))
	router.DELETE("/rooms/:id", middlewares.JWTAuth(handlers.DeleteRoomHandler))

	router.POST("/boards", middlewares.JWTAuth(handlers.CreateBoardHandler))
	router.GET("/boards", handlers.GetBoardsHandler)
	router.GET("/boards/:id", handlers.GetBoardByIDHandler)
	router.PUT("/boards/:id", middlewares.JWTAuth(handlers.UpdateBoardHandler))
	router.DELETE("/boards/:id", middlewares.JWTAuth(handlers.DeleteBoardHandler))

	router.POST("/games", middlewares.JWTAuth(handlers.CreateGameHandler))
	router.GET("/games", handlers.GetGamesHandler)
	router.GET("/games/:id", handlers.GetGameByIDHandler)
	router.DELETE("/games/:id", middlewares.JWTAuth(handlers.DeleteGameHandler))

	router.Handler("GET", "/swagger/*any", handlers.SwaggerUIHandler())

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
