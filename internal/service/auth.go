package service

import (
	"github.com/golang-jwt/jwt/v4"
	"homework/internal/entity"
	"homework/internal/infrastructure/repository"
	"time"
)

type AuthService struct {
	userRepo repository.UserRepository
	signKey  []byte
}

func NewAuthService(userRepo repository.UserRepository, signKey []byte) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		signKey:  []byte,
	}
}

type CustomerInfo struct {
	Name string
	Kind string
}

func (s *AuthService) createToken() (string, error) {

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	user := entity.User{"1", "test", "sadasd", "maga"}
	claims := entity.UserClaims{
		&jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
		"level1",
		user,
	}

	t.Claims = &claims

	return t.SignedString(s.signKey)
}
