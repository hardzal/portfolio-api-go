package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type WorkService interface {
	GetAllWork() ([]models.Work, error)
	GetWork(id string) (*models.Work, error)
	CreateWork(work *models.Work) (*models.Work, error)
	UpdateWork(work *models.Work) (*models.Work, error)
	DeleteWork(id string) error
}

type workService struct {
	workRepo repositories.WorkRepository
}

// CreateWork implements WorkService.
func (w *workService) CreateWork(work *models.Work) (*models.Work, error) {
	panic("unimplemented")
}

// DeleteWork implements WorkService.
func (w *workService) DeleteWork(id string) error {
	panic("unimplemented")
}

// GetAllWork implements WorkService.
func (w *workService) GetAllWork() ([]models.Work, error) {
	panic("unimplemented")
}

// GetWork implements WorkService.
func (w *workService) GetWork(id string) (*models.Work, error) {
	panic("unimplemented")
}

// UpdateWork implements WorkService.
func (w *workService) UpdateWork(work *models.Work) (*models.Work, error) {
	panic("unimplemented")
}

func NewWorkService(repo repositories.WorkRepository) WorkService {
	return &workService{workRepo: repo}
}
