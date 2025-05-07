package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type WorkService interface {
	GetAllWork() ([]models.Work, error)
	GetWork(id string) (*models.WorkResponse, error)
	CreateWork(work *models.WorkDTO, image *string) (*models.WorkResponse, error)
	UpdateWork(id string, work *models.WorkDTO, image *string) (*models.WorkResponse, error)
	DeleteWork(id string) error
}

type workService struct {
	workRepo repositories.WorkRepository
}

// CreateWork implements WorkService.
func (w *workService) CreateWork(work *models.WorkDTO, image *string) (*models.WorkResponse, error) {
	newWork := &models.Work{
		Role:        work.Role,
		Company:     work.Company,
		Description: work.Description,
		Stacks:      work.Stacks,
		Image:       *image,
		StartDate:   work.StartDate,
		EndDate:     work.EndDate,
	}

	newWorkCreate, err := w.workRepo.CreateWork(newWork)
	if err != nil {
		return nil, err
	}

	return w.mapToResponse(newWorkCreate), err
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
func (w *workService) UpdateWork(id string, work *models.WorkDTO, image *string) (*models.WorkResponse, error) {
	dataWork, err := w.workRepo.GetWork(id)
	if err != nil {
		return nil, err
	}

	newWork := &models.Work{
		ID:          dataWork.ID,
		Role:        work.Role,
		Company:     work.Company,
		Description: work.Description,
		Stacks:      work.Stacks,
		Image:       *image,
		StartDate:   work.StartDate,
		EndDate:     work.EndDate,
	}

	updateWork, err := w.workRepo.UpdateWork(newWork)
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
