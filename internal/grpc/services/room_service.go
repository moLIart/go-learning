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

type RoomService struct {
	generated.UnimplementedRoomServiceServer
}

func (s *RoomService) GetRoom(ctx context.Context, req *generated.RequestEntity) (*generated.GetRoomDto, error) {
	r, err := repository.GetRoomByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if r == nil {
		return nil, status.Errorf(codes.NotFound, "room not found")
	}
	return &generated.GetRoomDto{
		Id:   int32(r.ID),
		Code: r.Code,
	}, nil
}

func (s *RoomService) GetAllRooms(ctx context.Context, _ *emptypb.Empty) (*generated.RoomList, error) {
	rooms, err := repository.GetRooms()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	roomDtos := make([]*generated.GetRoomDto, len(rooms))
	for i, r := range rooms {
		roomDtos[i] = &generated.GetRoomDto{
			Id:   int32(r.ID),
			Code: r.Code,
		}
	}
	return &generated.RoomList{Rooms: roomDtos}, nil
}

func (s *RoomService) CreateRoom(ctx context.Context, req *generated.CreateRoomDto) (*generated.GetRoomDto, error) {
	r := room.NewRoom(req.Code)
	repository.AddEntity(r)
	return &generated.GetRoomDto{
		Id:   int32(r.ID),
		Code: r.Code,
	}, nil
}

func (s *RoomService) UpdateRoom(ctx context.Context, req *generated.UpdateRoomDto) (*generated.GetRoomDto, error) {
	ok, err := repository.UpdateRoomByID(int(req.Id), req.Code)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "room not found")
	}

	r, err := repository.GetRoomByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return &generated.GetRoomDto{
		Id:   int32(r.ID),
		Code: r.Code,
	}, nil
}

func (s *RoomService) DeleteRoom(ctx context.Context, req *generated.RequestEntity) (*emptypb.Empty, error) {
	ok, err := repository.DeleteRoomByID(int(req.Id))
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	if !ok {
		return nil, status.Errorf(codes.NotFound, "room not found")
	}
	return &emptypb.Empty{}, nil
}
