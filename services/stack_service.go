package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type StackService interface {
	GetAllStack() ([]models.Stack, error)
	GetStack(id string) (*models.Stack, error)
	CreateStack(stack *models.Stack) (*models.Stack, error)
	UpdateStack(stack *models.Stack) (*models.Stack, error)
	DeleteStack(id string) error
}

type stackService struct {
	stackRepo repositories.StackRepository
}

// CreateStack implements StackService.
func (s *stackService) CreateStack(stack *models.Stack) (*models.Stack, error) {
	panic("unimplemented")
}

// DeleteStack implements StackService.
func (s *stackService) DeleteStack(id string) error {
	panic("unimplemented")
}

// GetAllStack implements StackService.
func (s *stackService) GetAllStack() ([]models.Stack, error) {
	panic("unimplemented")
}

// GetStack implements StackService.
func (s *stackService) GetStack(id string) (*models.Stack, error) {
	panic("unimplemented")
}

// UpdateStack implements StackService.
func (s *stackService) UpdateStack(stack *models.Stack) (*models.Stack, error) {
	panic("unimplemented")
}

func NewStackService(repo repositories.StackRepository) StackService {
	return &stackService{stackRepo: repo}
}
