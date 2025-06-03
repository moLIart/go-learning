package services

import (
	"context"

	"github.com/moLIart/go-course/internal/grpc/generated"
	"github.com/moLIart/go-course/internal/model/room"
	"github.com/moLIart/go-course/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PlayerService struct {
	generated.UnimplementedPlayerServiceServer
}

func (s *PlayerService) GetPlayer(ctx context.Context, req *generated.RequestEntity) (*generated.GetPlayerDto, error) {
	player, err := repository.GetPlayerByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if player == nil {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}
	return &generated.GetPlayerDto{
		Id:   int32(player.ID),
		Name: player.Name,
	}, nil
}

func (s *PlayerService) GetAllPlayers(ctx context.Context, _ *emptypb.Empty) (*generated.PlayerList, error) {
	players, err := repository.GetPlayers()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	playerDtos := make([]*generated.GetPlayerDto, len(players))
	for i, player := range players {
		playerDtos[i] = &generated.GetPlayerDto{
			Id:   int32(player.ID),
			Name: player.Name,
		}
	}
	return &generated.PlayerList{Players: playerDtos}, nil
}

func (s *PlayerService) CreatePlayer(ctx context.Context, req *generated.CreatePlayerDto) (*generated.GetPlayerDto, error) {
	player := room.NewPlayer(req.Name)
	repository.AddEntity(player)
	return &generated.GetPlayerDto{
		Id:   int32(player.ID),
		Name: player.Name,
	}, nil
}

func (s *PlayerService) UpdatePlayer(ctx context.Context, req *generated.UpdatePlayerDto) (*generated.GetPlayerDto, error) {
	ok, err := repository.UpdatePlayerByID(int(req.Id), req.Name)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}
	player, err := repository.GetPlayerByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return &generated.GetPlayerDto{
		Id:   int32(player.ID),
		Name: player.Name,
	}, nil
}

func (s *PlayerService) DeletePlayer(ctx context.Context, req *generated.RequestEntity) (*emptypb.Empty, error) {
	ok, err := repository.DeletePlayerByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "player not found")
	}
	return &emptypb.Empty{}, nil
}
