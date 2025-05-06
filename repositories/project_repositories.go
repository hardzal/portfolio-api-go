package repositories

import (
	"errors"

	"github.com/hardzal/portfolio-api-go/models"
	"gorm.io/gorm"
)

var ErrProjectNotFound error = errors.New("project not found")

type ProjectRepository interface {
	GetProject(id string) (*models.Project, error)
	GetAllProjects() ([]models.Project, error)
	CreateProject(project *models.Project) (*models.Project, error)
	UpdateProject(project *models.Project) (*models.Project, error)
	DeleteProject(id string) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepository {
	return &projectRepository{db: db}
}

// CreateProject implements ProjectRepository.
func (p *projectRepository) CreateProject(project *models.Project) (*models.Project, error) {
	if err := p.db.Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

// DeleteProject implements ProjectRepository.
func (p *projectRepository) DeleteProject(id string) error {
	if err := p.db.Delete(&models.Project{}, id).Error; err != nil {
		return err
	}

	return nil
}

// GetProject implements ProjectRepository.
func (p *projectRepository) GetProject(id string) (*models.Project, error) {
	var project models.Project
	if err := p.db.First(&project, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrProjectNotFound
		}
		return nil, err
	}

	return &project, nil
}

// GetProjects implements ProjectRepository.
func (p *projectRepository) GetAllProjects() ([]models.Project, error) {
	var projects []models.Project
	if err := p.db.Find(&projects).Error; err != nil {
		return nil, err
	}

	return projects, nil
}

// UpdateProject implements ProjectRepository.
func (p *projectRepository) UpdateProject(project *models.Project) (*models.Project, error) {
	if err := p.db.Save(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}
