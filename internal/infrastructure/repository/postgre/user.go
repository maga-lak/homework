package postgre

import (
	"context"

	"github.com/go-pg/pg/v10"

	"homework/internal/entity"
)

type UserRepository struct {
	*pg.DB
}

func NewUserRepository(db *pg.DB) *UserRepository {
	return &UserRepository{
		db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *entity.User) error {
	return nil
}

func (r *UserRepository) GetUserByUserName(ctx context.Context, username string) (*entity.User, error) {
	//todo  не забыть username UNIQUE
	var user *entity.User
	err := r.DB.Model(&user).
		Where("username = ?", username).
		Context(ctx).
		First()
	if err != nil {
		return nil, err
	}

	return user, nil
}
