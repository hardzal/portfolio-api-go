package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

// inteface for abstract
type UserRepository interface {
	getUserByEmail(email string) (*models.User, error)
	getUserByUsername(username string) (*models.User, error)
}

// connect to db
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// getUserByEmail implements UserRepository.
func (u *userRepository) getUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(&models.User{Email: email}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}

// getUserByUsername implements UserRepository.
func (u *userRepository) getUserByUsername(username string) (*models.User, error) {
	var user models.User
	if err := u.db.Where(&models.User{Username: username}).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return &user, nil
}
