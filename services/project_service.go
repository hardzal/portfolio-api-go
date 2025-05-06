package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type ProjectService interface {
	GetAllProjects() ([]models.Project, error)
	GetProject(id string) (*models.Project, error)
	CreateProject(project *models.Project) (*models.Project, error)
	UpdateProject(project *models.Project) (*models.Project, error)
	DeleteProject(id string) error
}

type projectService struct {
	projectRepo repositories.ProjectRepository
}

// CreateProject implements ProjectService.
func (p *projectService) CreateProject(project *models.Project) (*models.Project, error) {
	panic("unimplemented")
}

// DeleteProject implements ProjectService.
func (p *projectService) DeleteProject(id string) error {
	panic("unimplemented")
}

// GetAllProjects implements ProjectService.
func (p *projectService) GetAllProjects() ([]models.Project, error) {
	panic("unimplemented")
}

// GetProject implements ProjectService.
func (p *projectService) GetProject(id string) (*models.Project, error) {
	panic("unimplemented")
}

// UpdateProject implements ProjectService.
func (p *projectService) UpdateProject(project *models.Project) (*models.Project, error) {
	panic("unimplemented")
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{projectRepo: repo}
}
