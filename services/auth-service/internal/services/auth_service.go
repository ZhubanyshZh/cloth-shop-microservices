package services

import (
	"errors"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/models"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/repositories"
	"github.com/ZhubanyshZh/cloth-shop-microservices/internal/utils"
)

type AuthService struct {
	Repo repositories.UserRepository
}

func NewAuthService(repo repositories.UserRepository) *AuthService {
	return &AuthService{Repo: repo}
}

func (s *AuthService) Register(email, password string) (*models.User, error) {
	_, err := s.Repo.FindByEmail(email)
	if err == nil {
		return nil, errors.New("user already exists")
	}
	hashed, _ := utils.HashPassword(password)
	return s.Repo.CreateUser(email, hashed)
}

func (s *AuthService) Login(email, password string) (access string, refresh string, err error) {
	user, err := s.Repo.FindByEmail(email)
	if err != nil {
		return "", "", errors.New("invalid credentials")
	}

	if !utils.CheckPasswordHash(password, user.PasswordHash) {
		return "", "", errors.New("invalid credentials")
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID.String())
	refreshToken, _ := utils.GenerateRefreshToken(user.ID.String())

	err = s.Repo.SaveRefreshToken(user.ID, refreshToken)
	if err != nil {
		return "", "", err
	}

	return accessToken, refreshToken, nil
}

func (s *AuthService) Logout(refreshToken string) error {
	return s.Repo.DeleteRefreshToken(refreshToken)
}
