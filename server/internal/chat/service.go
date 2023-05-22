package chat

import (
	"context"

	"github.com/danielwiratman/chatroom-with-media-uploads/util"
)

type ChatServiceImpl struct {
	roomRepo RoomRepository
}

func NewChatService(roomRepo RoomRepository) ChatService {
	return &ChatServiceImpl{roomRepo: roomRepo}
}

func (s *ChatServiceImpl) CreateRoom(ctx context.Context, dbtx util.DBTX, req CreateRoomReq) (*CreateRoomRes, error) {
	room, err := s.roomRepo.Create(ctx, dbtx, &Room{Name: req.Name})
	if err != nil {
		return nil, err
	}
	return &CreateRoomRes{ID: room.ID, Name: room.Name}, nil
}

func (s *ChatServiceImpl) GetRoomsByUserID(ctx context.Context, dbtx util.DBTX, userID int) ([]*Room, error) {
  rooms, err := s.roomRepo.GetByID(ctx, dbtx, userID)
  if err != nil {
    return nil, err
  }
  return rooms, nil
}
