package user

import (
	"context"
)

type RepositoryImpl struct{}

func NewRepository() Repository {
	return &RepositoryImpl{}
}

func (r *RepositoryImpl) Create(ctx context.Context, dbtx DBTX, user *User) (*User, error) {
	stmt, err := dbtx.PrepareContext(ctx, "INSERT INTO user_profile (name, username, email, password) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	id := 0
	err = stmt.QueryRowContext(ctx, user.Name, user.Username, user.Email, user.Password).Scan(&id)
	if err != nil {
		return nil, err
	}
	user.ID = int(id)
	return user, nil
}

func (r *RepositoryImpl) GetByEmail(ctx context.Context, dbtx DBTX, email string) (*User, error) {
	stmt, err := dbtx.PrepareContext(ctx, "SELECT id, name, username, email, password FROM user_profile WHERE email = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRowContext(ctx, email)
	user := &User{}
	err = row.Scan(&user.ID, &user.Name, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, nil
}
