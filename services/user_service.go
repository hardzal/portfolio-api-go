package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type UserService interface {
	GetUserByEmail(email string) (*models.UserResponse, error)
	GetUserByUsername(username string) (*models.UserResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{userRepo: repo}
}

func (u *userService) GetUserByEmail(email string) (*models.UserResponse, error) {
	user, err := u.userRepo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return u.mapToResponse(user), nil
}

func (u *userService) GetUserByUsername(username string) (*models.UserResponse, error) {
	user, err := u.userRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	return u.mapToResponse(user), nil
}

func (u *userService) mapToResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
