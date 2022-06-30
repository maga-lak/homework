package service

import (
	"context"
	"errors"
	"homework/internal/controller/http/v1/request"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"homework/internal/configs"
	"homework/internal/dto"
	"homework/internal/entity"
	"homework/internal/infrastructure/repository"
)

//todo отдельный файл? как в пыхе Exception
var (
	UserFoundErr       = errors.New("user not found")
	InvalidPasswordErr = errors.New("invalid password")
)

type AuthService struct {
	userRepo repository.UserRepository
	config   configs.AuthConfig
}

func NewAuthService(userRepo repository.UserRepository, config configs.AuthConfig) *AuthService {
	return &AuthService{
		userRepo: userRepo,
		config:   config,
	}
}

func (s *AuthService) Authorize(ctx context.Context, data *request.AuthRequest) (*dto.AuthView, error) {
	user, err := s.userRepo.GetUserByUserName(ctx, data.Username)
	if err != nil {
		return nil, UserFoundErr
	}

	if !CheckPasswordHash(data.Password, user.Password) {
		return nil, InvalidPasswordErr
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

func (s *AuthService) createToken(user entity.User) (string, error) {
	t := jwt.New(jwt.GetSigningMethod(s.config.Alg))

	claims := entity.UserClaims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.config.ExpireDuration)),
		},
		User: user,
	}

	t.Claims = &claims

	return t.SignedString(s.config.SignKey)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
