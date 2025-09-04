package service

import (
	"errors"
	model "gin-basic/internal/models"
	"gin-basic/internal/repository"
)

type AuthService struct {
	repo *repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) *AuthService {
	return &AuthService{repo: &repo}
}

func (s *AuthService) LoginService(email string, pass string) (*model.UserModel, error) {

	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}

	if user.Password != pass {
		return nil, errors.New("password incorrect")
	}

	return user, nil
}
