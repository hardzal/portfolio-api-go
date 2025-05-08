package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type StackService interface {
	GetAllStack() ([]models.Stack, error)
	GetStack(id uint) (*models.StackResponse, error)
	CreateStack(stack *models.StackDTO, image *string) (*models.StackResponse, error)
	UpdateStack(id uint, stack *models.StackDTO, updatedImage string) (*models.StackResponse, error)
	DeleteStack(id uint) error
}

type stackService struct {
	stackRepo repositories.StackRepository
}

// CreateStack implements StackService.
func (s *stackService) CreateStack(stack *models.StackDTO, image *string) (*models.StackResponse, error) {
	newStack := &models.Stack{
		Name:  stack.Name,
		Image: *image,
	}

	newStack, err := s.stackRepo.CreateStack(newStack)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(newStack), err
}

// DeleteStack implements StackService.
func (s *stackService) DeleteStack(id uint) error {
	if err := s.stackRepo.DeleteStack(id); err != nil {
		return err
	}

	return nil
}

// GetAllStack implements StackService.
func (s *stackService) GetAllStack() ([]models.Stack, error) {
	stacks, err := s.stackRepo.GetAllStacks()
	if err != nil {
		return nil, err
	}

	return stacks, nil
}

// GetStack implements StackService.
func (s *stackService) GetStack(id uint) (*models.StackResponse, error) {
	stack, err := s.stackRepo.GetStack(id)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(stack), err
}

// UpdateStack implements StackService.
func (s *stackService) UpdateStack(id uint, stack *models.StackDTO, updatedImage string) (*models.StackResponse, error) {
	dataStack, err := s.stackRepo.GetStack(id)
	if err != nil {
		return nil, err
	}

	if stack.Name != "" {
		dataStack.Name = stack.Name
	}

	if updatedImage != "" {
		dataStack.Image = updatedImage
	}

	updateddStack, err := s.stackRepo.UpdateStack(dataStack)
	if err != nil {
		return nil, err
	}

	return s.mapToResponse(updateddStack), err
}

func NewStackService(repo repositories.StackRepository) StackService {
	return &stackService{stackRepo: repo}
}

func (s *stackService) mapToResponse(stack *models.Stack) *models.StackResponse {
	return &models.StackResponse{
		ID:        stack.ID,
		Name:      stack.Name,
		Image:     stack.Image,
		CreatedAt: stack.CreatedAt,
		UpdatedAt: stack.UpdatedAt,
	}
}
