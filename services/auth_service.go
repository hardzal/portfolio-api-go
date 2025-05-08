package services

import (
	"errors"
	"fmt"

	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
	jwt "github.com/hardzal/portfolio-api-go/utils/jwt"
	password "github.com/hardzal/portfolio-api-go/utils/password"
)

var ErrInvalidCredentials = errors.New("invalid email or password")

type AuthService interface {
	LoginUser(user models.UserLoginDTO) (string, error)
}

type authService struct {
	userAuth repositories.UserRepository
}

func NewAuthService(user repositories.UserRepository) AuthService {
	return &authService{userAuth: user}
}

func (a *authService) LoginUser(user models.UserLoginDTO) (string, error) {
	userLogin, err := a.userAuth.GetUserByEmail(user.Email)

	if err != nil {
		if errors.Is(err, repositories.ErrUserNotFound) {
			return "", ErrInvalidCredentials
		}

		return "", fmt.Errorf("auth.Login: %w", err)
	}

	if err := password.Verify(userLogin.Password, user.Password); err != nil {
		return "", ErrInvalidCredentials
	}

	t := jwt.Generate(&jwt.TokenPayload{
		ID:       userLogin.ID.String(),
		Username: userLogin.Username,
	})

	return t, nil
}
