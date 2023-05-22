package chat

import (
	"context"

	"github.com/danielwiratman/chatroom-with-media-uploads/util"
)

type RoomRepositoryImpl struct{}

func NewRoomRepository() RoomRepository {
	return &RoomRepositoryImpl{}
}

func (r *RoomRepositoryImpl) Create(ctx context.Context, dbtx util.DBTX, room *Room) (*Room, error) {
	stmt, err := dbtx.PrepareContext(ctx, "INSERT INTO room (name) VALUES ($1) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	id := 0
	err = stmt.QueryRowContext(ctx, room.Name).Scan(&id)
	if err != nil {
		return nil, err
	}
	room.ID = int(id)
	return room, nil
}

func (r *RoomRepositoryImpl) GetByID(ctx context.Context, dbtx util.DBTX, id int) (*Room, error) {
	stmt, err := dbtx.PrepareContext(ctx, "SELECT id, name FROM room WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, id)
	room := &Room{}
	err = row.Scan(&room.ID, &room.Name)
	if err != nil {
		return nil, err
	}
	return room, nil
}

func (r *RoomRepositoryImpl) Delete(ctx context.Context, dbtx util.DBTX, id int) error {
	stmt, err := dbtx.PrepareContext(ctx, "DELETE FROM room WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	return err

}
