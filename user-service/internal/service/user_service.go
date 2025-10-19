package service

import (
	"fmt"
	"go-ms-demo/user-service/internal/domain"
)

type UserService struct {
	userRepo domain.UserRepository
}

func NewUserService(userRepo domain.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) CreateUser(u *domain.User) (*domain.User, error) {
	user, err := s.userRepo.InsertUser(u)
	if err != nil {
		return nil, fmt.Errorf("could not create user: %w", err)
	}
	return user, nil
}

func (s *UserService) GetAllUsers() ([]domain.User, error) {
	users, err := s.userRepo.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("could not retrieve all users: %w", err)
	}
	return users, nil
}

func (s *UserService) GetUser(id int64) (*domain.User, error) {
	user, err := s.userRepo.GetUser(id)
	if err != nil {
		return nil, fmt.Errorf("could not retrieve user: %w", err)
	}
	return user, nil
}

func (s *UserService) UpdateUser(u *domain.User) (*domain.User, error) {
	user, err := s.userRepo.UpdateUser(u)
	if err != nil {
		return nil, fmt.Errorf("could not update user: %w", err)
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int64) (int64, error) {
	rowsAffected, err := s.userRepo.DeleteUser(id)
	if err != nil {
		return 0, fmt.Errorf("could not delete user: %w", err)
	}
	return rowsAffected, nil
}
