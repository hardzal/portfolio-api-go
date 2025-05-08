package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

var ErrAboutNotFound error = errors.New("about not fund")

type AboutRepository interface {
	GetAbout(id uint) (*models.About, error)
	CreateAbout(about *models.About) (*models.About, error)
	UpdateAbout(about *models.About) (*models.About, error)
}

type aboutRepository struct {
	db *gorm.DB
}

func NewAboutRepository(db *gorm.DB) AboutRepository {
	return &aboutRepository{db: db}
}

// CreateAbout implements AboutRepository.
func (a *aboutRepository) CreateAbout(about *models.About) (*models.About, error) {
	if err := a.db.Create(about).Error; err != nil {
		return nil, err
	}

	return about, nil
}

// GetAbout implements AboutRepository.
func (a *aboutRepository) GetAbout(id uint) (*models.About, error) {
	var about models.About
	if err := a.db.First(&about, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrAboutNotFound
		}

		return nil, err
	}

	return &about, nil
}

// UpdateAbout implements AboutRepository.
func (a *aboutRepository) UpdateAbout(about *models.About) (*models.About, error) {
	if err := a.db.Save(about).Error; err != nil {
		return nil, err
	}

	return about, nil
}
