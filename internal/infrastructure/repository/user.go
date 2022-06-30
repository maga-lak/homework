package repository

import (
	"context"
	"homework/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *entity.User) error
	GetUser(ctx context.Context, username, password string) (*entity.User, error)
}
