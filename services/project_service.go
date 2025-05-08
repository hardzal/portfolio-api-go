package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
	"github.com/lib/pq"
)

type ProjectService interface {
	GetAllProjects() ([]models.Project, error)
	GetProject(id uint) (*models.ProjectResponse, error)
	CreateProject(project *models.ProjectDTO, image *string) (*models.ProjectResponse, error)
	UpdateProject(id uint, project *models.ProjectDTO, image *string) (*models.ProjectResponse, error)
	DeleteProject(id uint) error
}

type projectService struct {
	projectRepo repositories.ProjectRepository
}

// CreateProject implements ProjectService.
func (p *projectService) CreateProject(project *models.ProjectDTO, newImage *string) (*models.ProjectResponse, error) {

	newProject := &models.Project{
		Title:       project.Title,
		Description: project.Description,
		Stacks:      pq.StringArray(project.Stacks),
		ImageUrl:    newImage,
		Repo:        project.Repo,
		Demo:        project.Demo,
	}

	newProjectCreate, err := p.projectRepo.CreateProject(newProject)
	if err != nil {
		return nil, err
	}

	return p.mapToResponse(newProjectCreate), err
}

// DeleteProject implements ProjectService.
func (p *projectService) DeleteProject(id uint) error {
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
func (p *projectService) GetProject(id uint) (*models.ProjectResponse, error) {
	project, err := p.projectRepo.GetProject(id)
	if err != nil {
		return nil, err
	}

	return p.mapToResponse(project), err
}

// UpdateProject implements ProjectService.
func (p *projectService) UpdateProject(id uint, project *models.ProjectDTO, updatedImage *string) (*models.ProjectResponse, error) {
	dataProject, err := p.projectRepo.GetProject(id)
	if err != nil {
		return nil, err
	}

	if project.Title != "" {
		dataProject.Title = project.Title
	}

	if project.Description != "" {
		dataProject.Description = project.Description
	}

	if project.Demo != nil {
		dataProject.Demo = project.Demo
	}

	if project.Repo != nil {
		dataProject.Repo = project.Repo
	}

	if len(project.Stacks) != 0 || project.Stacks != nil {
		dataProject.Description = project.Description
	}

	if updatedImage != nil {
		dataProject.ImageUrl = updatedImage
	}

	resProject, err := p.projectRepo.UpdateProject(dataProject)
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
