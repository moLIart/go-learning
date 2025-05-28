package main

import (
	"context"
	"log"

	"github.com/moLIart/go-course/internal/grpc/generated"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:8081", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer conn.Close()

	client := generated.NewBoardServiceClient(conn)

	res, err := client.CreateBoard(context.Background(), &generated.CreateBoardDto{Size: 5})
	if err != nil {
		log.Fatalf("Error creating board: %v", err)
	}
	log.Printf("Board created with ID: %d", res.Id)

	resGet, err := client.GetBoard(context.Background(), &generated.RequestEntity{Id: res.Id})
	if err != nil {
		log.Fatalf("Error getting board: %v", err)
	}
	log.Printf("Retrieved board with ID: %d, Size: %d", resGet.Id, resGet.Size)
}
