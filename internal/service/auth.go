package service

import (
	"context"
	"crypto/sha1"
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"

	"homework/internal/dto"
	"homework/internal/entity"
	"homework/internal/infrastructure/repository"
)

// todo configs
const (
	alg      = "RS256"
	hashSalt = "dnfldfg"
)

//todo отдельный файл? как в пыхе Exception
var UserNotFoundErr = errors.New("user not found")

type AuthService struct {
	userRepo repository.UserRepository
	signKey  []byte
	alg      string
}

func NewAuthService(userRepo repository.UserRepository, signKey []byte, alg string) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		signKey:  signKey,
		alg:      alg,
	}
}

func (s *AuthService) Authorize(ctx context.Context, username, password string) (*dto.AuthView, error) {
	password = s.getPasswordHash(password)

	user, err := s.userRepo.GetUser(ctx, username, password)
	if err != nil {
		return nil, UserNotFoundErr
	}

	token, err := s.createToken(*user)
	if err != nil {
		return nil, err
	}

	view := dto.AuthView{
		Token:    token,
		UserId:   user.ID,
		UserName: user.Name,
	}

	return &view, nil
}

func (s *AuthService) getPasswordHash(password string) string {
	pwd := sha1.New()
	pwd.Write([]byte(password))
	pwd.Write([]byte(hashSalt))
	password = fmt.Sprintf("%x", pwd.Sum(nil))
	return password
}

func (s *AuthService) createToken(user entity.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(s.alg))

	claims := entity.UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * 1)),
		},
		User: user,
	}

	t.Claims = &claims

	return t.SignedString(s.signKey)
}
