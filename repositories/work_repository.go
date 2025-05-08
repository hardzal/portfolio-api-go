package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

var ErrWorkNotFound error = errors.New("work not found")

type WorkRepository interface {
	GetAllWorks() ([]models.Work, error)
	GetWork(id uint) (*models.Work, error)
	CreateWork(work *models.Work) (*models.Work, error)
	UpdateWork(work *models.Work) (*models.Work, error)
	DeleteWork(id uint) error
}

type workRepository struct {
	db *gorm.DB
}

func NewWorkRepository(db *gorm.DB) WorkRepository {
	return &workRepository{db: db}
}

// CreateWork implements WorkRepository.
func (w *workRepository) CreateWork(work *models.Work) (*models.Work, error) {
	if err := w.db.Create(work).Error; err != nil {
		return nil, err
	}

	return work, nil
}

// DeleteWork implements WorkRepository.
func (w *workRepository) DeleteWork(id uint) error {
	if err := w.db.Delete(&models.Work{}, id).Error; err != nil {
		return err
	}

	return nil
}

// GetAllWorks implements WorkRepository.
func (w *workRepository) GetAllWorks() ([]models.Work, error) {
	var works []models.Work
	if err := w.db.Find(&works).Error; err != nil {
		return nil, err
	}

	return works, nil
}

// GetWork implements WorkRepository.
func (w *workRepository) GetWork(id uint) (*models.Work, error) {
	var work models.Work
	if err := w.db.First(&work, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrWorkNotFound
		}

		return nil, err
	}

	return &work, nil
}

// UpdateWork implements WorkRepository.
func (w *workRepository) UpdateWork(work *models.Work) (*models.Work, error) {
	if err := w.db.Save(work).Error; err != nil {
		return nil, err
	}

	return work, nil
}
