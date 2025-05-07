package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type ProjectService interface {
	GetAllProjects() ([]models.Project, error)
	GetProject(id string) (*models.ProjectResponse, error)
	CreateProject(project *models.ProjectDTO, image *string) (*models.ProjectResponse, error)
	UpdateProject(id string, project *models.ProjectDTO, image *string) (*models.ProjectResponse, error)
	DeleteProject(id string) error
}

type projectService struct {
	projectRepo repositories.ProjectRepository
}

// CreateProject implements ProjectService.
func (p *projectService) CreateProject(project *models.ProjectDTO, newImage *string) (*models.ProjectResponse, error) {
	newProject := &models.Project{
		Title:       project.Title,
		Description: project.Description,
		Stacks:      project.Stacks,
		ImageUrl:    newImage,
		Repo:        project.Repo,
		Demo:        project.Demo,
	}

	newProject, err := p.projectRepo.CreateProject(newProject)
	if err != nil {
		return nil, err
	}

	return p.mapToResponse(newProject), err
}

// DeleteProject implements ProjectService.
func (p *projectService) DeleteProject(id string) error {
	if err := p.projectRepo.DeleteProject(id); err != nil {
		return err
	}

	return nil
}

// GetAllProjects implements ProjectService.
func (p *projectService) GetAllProjects() ([]models.Project, error) {
	projects, err := p.projectRepo.GetAllProjects()
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// GetProject implements ProjectService.
func (p *projectService) GetProject(id string) (*models.ProjectResponse, error) {
	project, err := p.projectRepo.GetProject(id)
	if err != nil {
		return nil, err
	}

	return p.mapToResponse(project), err
}

// UpdateProject implements ProjectService.
func (p *projectService) UpdateProject(id string, project *models.ProjectDTO, updatedImage *string) (*models.ProjectResponse, error) {
	dataProject, err := p.projectRepo.GetProject(id)
	if err != nil {
		return nil, err
	}

	newProject := &models.Project{
		ID:          dataProject.ID,
		Title:       project.Title,
		Description: project.Description,
		Stacks:      project.Stacks,
		ImageUrl:    updatedImage,
		Repo:        project.Repo,
		Demo:        project.Demo,
	}

	resProject, err := p.projectRepo.UpdateProject(newProject)
	if err != nil {
		return nil, err
	}

	return p.mapToResponse(resProject), err
}

func NewProjectService(repo repositories.ProjectRepository) ProjectService {
	return &projectService{projectRepo: repo}
}

func (p *projectService) mapToResponse(project *models.Project) *models.ProjectResponse {
	return &models.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		ImageUrl:    project.ImageUrl,
		Stacks:      project.Stacks,
		Repo:        project.Repo,
		Demo:        project.Demo,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}
