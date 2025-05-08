package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
	"github.com/lib/pq"
)

type WorkService interface {
	GetAllWork() ([]models.Work, error)
	GetWork(id uint) (*models.WorkResponse, error)
	CreateWork(work *models.WorkDTO, image *string) (*models.WorkResponse, error)
	UpdateWork(id uint, work *models.WorkDTO, image *string) (*models.WorkResponse, error)
	DeleteWork(id uint) error
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
		Stacks:      pq.StringArray(work.Stacks),
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
func (w *workService) DeleteWork(id uint) error {
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
func (w *workService) GetWork(id uint) (*models.WorkResponse, error) {
	work, err := w.workRepo.GetWork(id)
	if err != nil {
		return nil, err
	}

	return w.mapToResponse(work), err
}

// UpdateWork implements WorkService.
func (w *workService) UpdateWork(id uint, work *models.WorkDTO, image *string) (*models.WorkResponse, error) {
	dataWork, err := w.workRepo.GetWork(id)
	if err != nil {
		return nil, err
	}

	if work.Role != "" {
		dataWork.Role = work.Role
	}

	if work.Company != "" {
		dataWork.Company = work.Company
	}

	if work.Description != "" {
		dataWork.Description = work.Description
	}

	if len(work.Stacks) != 0 {
		dataWork.Stacks = pq.StringArray(work.Stacks)
	}

	if image != nil {
		dataWork.Image = *image
	}

	if work.StartDate != nil {
		dataWork.StartDate = work.StartDate
	}

	if work.EndDate != nil {
		dataWork.EndDate = work.EndDate
	}

	updateWork, err := w.workRepo.UpdateWork(dataWork)
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
