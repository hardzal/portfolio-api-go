package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

var ErrUserNotFound = errors.New("user not found")

// inteface for abstract
type UserRepository interface {
	GetUserByEmail(email string) (*models.User, error)
	GetUserByUsername(username string) (*models.User, error)
}

// connect to db
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// getUserByEmail implements UserRepository.
func (u *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return &user, nil
}

// getUserByUsername implements UserRepository.
func (u *userRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
