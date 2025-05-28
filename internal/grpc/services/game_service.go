package services

import (
	"context"

	"github.com/moLIart/go-course/internal/grpc/generated"
	"github.com/moLIart/go-course/internal/model/game"
	"github.com/moLIart/go-course/internal/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GameService struct {
	generated.UnimplementedGameServiceServer
}

func (s *GameService) GetGame(ctx context.Context, req *generated.RequestEntity) (*generated.GetGameDto, error) {
	g := repository.GetGameByID(int(req.Id))
	if g == nil {
		return nil, status.Errorf(codes.NotFound, "game not found")
	}
	return &generated.GetGameDto{
		Id: int32(g.ID),
	}, nil
}

func (s *GameService) GetAllGames(ctx context.Context, _ *emptypb.Empty) (*generated.GameList, error) {
	games := repository.GetGames()
	gameDtos := make([]*generated.GetGameDto, len(games))
	for i, g := range games {
		gameDtos[i] = &generated.GetGameDto{
			Id: int32(g.ID),
		}
	}
	return &generated.GameList{Games: gameDtos}, nil
}

func (s *GameService) CreateGame(ctx context.Context, _ *emptypb.Empty) (*generated.GetGameDto, error) {
	g := game.NewGame()
	repository.AddEntity(g)
	return &generated.GetGameDto{
		Id: int32(g.ID),
	}, nil
}

func (s *GameService) DeleteGame(ctx context.Context, req *generated.RequestEntity) (*emptypb.Empty, error) {
	ok := repository.DeleteGameByID(int(req.Id))
	if !ok {
		return nil, status.Errorf(codes.NotFound, "game not found")
	}
	return &emptypb.Empty{}, nil
}
