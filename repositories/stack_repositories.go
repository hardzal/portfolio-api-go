package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

var ErrStackNotFound error = errors.New("stack not found")

type StackRepository interface {
	GetAllStacks() ([]models.Stack, error)
	GetStack(id string) (*models.Stack, error)
	CreateStack(stack *models.Stack) (*models.Stack, error)
	UpdateProject(stack *models.Stack) (*models.Stack, error)
	DeleteStack(id string) error
}

type stackRepository struct {
	db *gorm.DB
}

func NewStackRepository(db *gorm.DB) StackRepository {
	return &stackRepository{db: db}
}

// CreateStack implements StackRepository.
func (s *stackRepository) CreateStack(stack *models.Stack) (*models.Stack, error) {
	if err := s.db.Create(stack).Error; err != nil {
		return nil, err
	}

	return stack, nil
}

// DeleteStack implements StackRepository.
func (s *stackRepository) DeleteStack(id string) error {
	if err := s.db.Delete(&models.Stack{}, id).Error; err != nil {
		return err
	}

	return nil
}

// GetAllStacks implements StackRepository.
func (s *stackRepository) GetAllStacks() ([]models.Stack, error) {
	var stacks []models.Stack
	if err := s.db.Find(&stacks).Error; err != nil {
		return nil, err
	}

	return stacks, nil
}

// GetStack implements StackRepository.
func (s *stackRepository) GetStack(id string) (*models.Stack, error) {
	var stack models.Stack
	if err := s.db.First(&stack, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrStackNotFound
		}

		return nil, err
	}

	return &stack, nil
}

// UpdateProject implements StackRepository.
func (s *stackRepository) UpdateProject(stack *models.Stack) (*models.Stack, error) {
	if err := s.db.Save(stack).Error; err != nil {
		return nil, err
	}

	return stack, nil
}
