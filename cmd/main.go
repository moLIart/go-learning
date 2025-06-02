package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/moLIart/go-course/internal"
	"github.com/moLIart/go-course/internal/grpc/generated"
	"github.com/moLIart/go-course/internal/grpc/services"
	"github.com/moLIart/go-course/internal/middlewares"
	"github.com/moLIart/go-course/internal/repository"
	"google.golang.org/grpc"
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

	middlewares.SetJwtSecret(os.Getenv("JWT_SECRET"))

	repository.Startup(os.Getenv("MONGO_DS"), os.Getenv("REDIS_DS"))

	httpSrv := &http.Server{
		Addr:    ":8080",
		Handler: internal.RegisterHTTPRoutes(),
	}

	grpcSvc := grpc.NewServer()
	generated.RegisterBoardServiceServer(grpcSvc, &services.BoardService{})
	generated.RegisterGameServiceServer(grpcSvc, &services.GameService{})
	generated.RegisterPlayerServiceServer(grpcSvc, &services.PlayerService{})
	generated.RegisterRoomServiceServer(grpcSvc, &services.RoomService{})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-signalChan

		log.Println("Shutting down HTTP server...")
		if err := httpSrv.Shutdown(context.Background()); err != nil {
			log.Fatalf("HTTP Server forced to shutdown: %s", err)
		}

		log.Println("Shutting down gRPC server...")
		grpcSvc.GracefulStop()
	}()

	go func() {
		log.Println("Starting HTTP server on :8080")
		if err := httpSrv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("Could not start HTTP server: %s", err)
		}
	}()

	go func() {
		log.Println("Starting gRPC server on :8081")
		lis, err := net.Listen("tcp", ":8081")
		if err != nil {
			log.Fatalf("Failed to listen: %v", err)
		}

		if err := grpcSvc.Serve(lis); err != nil {
			log.Fatalf("Failed to serve gRPC server: %v", err)
		}
	}()

	wg.Wait()
	log.Println("Server gracefully stopped")
}
