package db

import (
	"database/sql"
  _ "github.com/lib/pq"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:admin@localhost:5432/chatroom?sslmode=disable")
  if err != nil {
    panic(err)
  }
  db.SetConnMaxLifetime(time.Hour)
  db.SetConnMaxIdleTime(time.Minute * 5)
  db.SetMaxOpenConns(10)
  db.SetMaxIdleConns(10)
  return db
}
