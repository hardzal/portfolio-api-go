package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type WorkService interface {
	GetAllWork() ([]models.Work, error)
	GetWork(id string) (*models.WorkResponse, error)
	CreateWork(work *models.Work) (*models.WorkResponse, error)
	UpdateWork(work *models.Work) (*models.WorkResponse, error)
	DeleteWork(id string) error
}

type workService struct {
	workRepo repositories.WorkRepository
}

// CreateWork implements WorkService.
func (w *workService) CreateWork(work *models.Work) (*models.WorkResponse, error) {
	newWork, err := w.workRepo.CreateWork(work)
	if err != nil {
		return nil, err
	}

	return w.mapToResponse(newWork), err
}

// DeleteWork implements WorkService.
func (w *workService) DeleteWork(id string) error {
	if err := w.workRepo.DeleteWork(id); err != nil {
		return err
	}

	return nil
}

// GetAllWork implements WorkService.
func (w *workService) GetAllWork() ([]models.Work, error) {
	works, err := w.workRepo.GetAllWorks()
	if err != nil {
		return nil, err
	}

	return works, err
}

// GetWork implements WorkService.
func (w *workService) GetWork(id string) (*models.WorkResponse, error) {
	work, err := w.workRepo.GetWork(id)
	if err != nil {
		return nil, err
	}

	return w.mapToResponse(work), err
}

// UpdateWork implements WorkService.
func (w *workService) UpdateWork(work *models.Work) (*models.WorkResponse, error) {
	updateWork, err := w.workRepo.UpdateWork(work)
	if err != nil {
		return nil, err
	}

	return w.mapToResponse(updateWork), err
}

func NewWorkService(repo repositories.WorkRepository) WorkService {
	return &workService{workRepo: repo}
}

func (w *workService) mapToResponse(work *models.Work) *models.WorkResponse {
	return &models.WorkResponse{
		ID:          work.ID,
		Role:        work.Role,
		Company:     work.Company,
		Description: work.Description,
		Stacks:      work.Stacks,
		Image:       work.Image,
		StartDate:   work.StartDate,
		EndDate:     work.EndDate,
		CreatedAt:   work.CreatedAt,
		UpdatedAt:   work.UpdatedAt,
	}
}
