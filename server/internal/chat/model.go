package chat

import (
	"context"
	"time"

	"github.com/danielwiratman/chatroom-with-media-uploads/util"
)

type Room struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type RoomRepository interface {
	Create(ctx context.Context, dbtx util.DBTX, room *Room) (*Room, error)
	GetByID(ctx context.Context, dbtx util.DBTX, id int) (*Room, error)
	Delete(ctx context.Context, dbtx util.DBTX, id int) error
}

type CreateRoomReq struct {
	Name string `json:"name"`
}

type CreateRoomRes struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	ID        int       `json:"id"`
	Text      string    `json:"text"`
	UserID    int       `json:"user_id"`
	RoomID    int       `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRoom struct {
	ID        int       `json:"id"`
	UserID    int       `json:"user_id"`
	RoomID    int       `json:"room_id"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRoomRepository interface {
	Create(ctx context.Context, dbtx util.DBTX, userRoom *UserRoom) (*UserRoom, error)
}

type ChatService interface {
	CreateRoom(ctx context.Context, dbtx util.DBTX, req CreateRoomReq) (*CreateRoomRes, error)
	GetRoomsByUserID(ctx context.Context, dbtx util.DBTX, userID int) ([]*Room, error)
}
