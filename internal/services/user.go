package services

import (
	"github.com/salasi00/go-market/internal/models"
	"github.com/salasi00/go-market/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (s *UserService) CreateUser(user models.User) (models.User, error) {
	createdUser, err := s.userRepository.CreateUser(user)
	if err != nil {
		return models.User{}, err
	}

	return createdUser, nil
}

func (s *UserService) GetUser(ID int) (models.User, error) {
	user, err := s.userRepository.GetUser(ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
