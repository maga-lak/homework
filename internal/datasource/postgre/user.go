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

func (r *UserRepository) GetUser(ctx context.Context, username, password string) (*entity.User, error) {
	var user *entity.User
	//хэш пароля
	err := r.DB.Model(&user).
		Where("username = ? and password= ?", username, password).
		Context(ctx).
		First()
	if err != nil {
		return nil, err
	}

	return user, nil
}
