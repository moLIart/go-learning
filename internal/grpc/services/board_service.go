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

type BoardService struct {
	generated.UnimplementedBoardServiceServer
}

func (s *BoardService) GetBoard(ctx context.Context, req *generated.RequestEntity) (*generated.GetBoardDto, error) {
	board, err := repository.GetBoardByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if board == nil {
		return nil, status.Errorf(codes.NotFound, "board not found")
	}
	return &generated.GetBoardDto{
		Id:   int32(board.ID),
		Size: int32(board.Size),
	}, nil
}

func (s *BoardService) GetAllBoards(ctx context.Context, _ *emptypb.Empty) (*generated.BoardList, error) {
	boards, err := repository.GetBoards()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	boardDtos := make([]*generated.GetBoardDto, len(boards))
	for i, board := range boards {
		boardDtos[i] = &generated.GetBoardDto{
			Id:   int32(board.ID),
			Size: int32(board.Size),
		}
	}
	return &generated.BoardList{Boards: boardDtos}, nil
}

func (s *BoardService) CreateBoard(ctx context.Context, req *generated.CreateBoardDto) (*generated.GetBoardDto, error) {
	board := game.NewBoard(int(req.Size))
	repository.AddEntity(board)
	return &generated.GetBoardDto{
		Id:   int32(board.ID),
		Size: int32(board.Size),
	}, nil
}

func (s *BoardService) UpdateBoard(ctx context.Context, req *generated.UpdateBoardDto) (*generated.GetBoardDto, error) {
	ok, err := repository.UpdateBoardByID(int(req.Id), int(req.Size))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "board not found")
	}

	board, err := repository.GetBoardByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return &generated.GetBoardDto{
		Id:   int32(board.ID),
		Size: int32(board.Size),
	}, nil
}

func (s *BoardService) DeleteBoard(ctx context.Context, req *generated.RequestEntity) (*emptypb.Empty, error) {
	ok, err := repository.DeleteBoardByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "board not found")
	}
	return &emptypb.Empty{}, nil
}
