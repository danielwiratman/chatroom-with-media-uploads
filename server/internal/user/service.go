package user

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/danielwiratman/chatroom-with-media-uploads/util"
)

type ServiceImpl struct {
	repo Repository
	db   *sql.DB
}

func NewService(repo Repository, db *sql.DB) Service {
	return &ServiceImpl{repo: repo, db: db}
}

func (s *ServiceImpl) Create(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
  if (req.Username == "") || (req.Email == "") || (req.Password == "") {
    return nil, errors.New("name, username, email and password are required") 
  }
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	user, err := s.repo.Create(ctx, tx, &User{
		Name:     req.Name,
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	})
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	return &CreateUserRes{
		ID:       strconv.Itoa(user.ID),
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
	}, nil
}

func (s *ServiceImpl) Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	tx, err := s.db.BeginTx(ctx, nil)
	user, err := s.repo.GetByEmail(ctx, tx, req.Email)
	if err != nil {
		return nil, err
	}
	err = tx.Commit()
	if err != nil {
		return nil, err
	}
	if err = util.CheckPassword(req.Password, user.Password); err != nil {
		return nil, err
	}
	jwtToken, err := util.GenerateJWT(user.ID)
	if err != nil {
		return nil, err
	}
	return &LoginUserRes{
		ID:       strconv.Itoa(user.ID),
		Name:     user.Name,
		Username: user.Username,
		Email:    user.Email,
		Token:    jwtToken,
	}, nil
}
