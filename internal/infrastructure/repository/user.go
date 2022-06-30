package repository

import (
	"context"
	"homework/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUserByUserName(ctx context.Context, username string) (*entity.User, error)
}
