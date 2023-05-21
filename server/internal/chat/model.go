package chat

import (
	"time"
)

type Room struct {
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

type Repository interface {
}

type Service interface {
}
