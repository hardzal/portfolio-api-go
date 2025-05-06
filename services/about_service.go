package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type AboutService interface {
	GetAbout(id string) (*models.AboutResponse, error)
	CreateAbout(about models.About) (*models.AboutResponse, error)
	UpdateAbout(about *models.About) (*models.AboutResponse, error)
}

type aboutService struct {
	aboutRepo repositories.AboutRepository
}

func NewAboutService(repo repositories.AboutRepository) AboutService {
	return &aboutService{aboutRepo: repo}
}

// CreateAbout implements AboutService.
func (a *aboutService) CreateAbout(about models.About) (*models.AboutResponse, error) {

	newAbout, err := a.aboutRepo.CreateAbout(&about)
	if err != nil {
		return nil, err
	}

	return a.mapToResponse(newAbout), nil
}

// GetAbout implements AboutService.
func (a *aboutService) GetAbout(id string) (*models.AboutResponse, error) {
	about, err := a.aboutRepo.GetAbout(id)
	if err != nil {
		return nil, err
	}

	return a.mapToResponse(about), nil
}

// UpdateAbout implements AboutService.
func (a *aboutService) UpdateAbout(about *models.About) (*models.AboutResponse, error) {
	about, err := a.aboutRepo.UpdateAbout(about)

	if err != nil {
		return nil, err
	}

	return a.mapToResponse(about), nil
}

func (a *aboutService) mapToResponse(about *models.About) *models.AboutResponse {
	return &models.AboutResponse{
		ID:          about.ID,
		Title:       about.Title,
		Profession:  about.Profession,
		Description: about.Description,
		Location:    about.Location,
		IsAvailable: about.IsAvailable,
		Handphone:   about.Handphone,
		Email:       about.Email,
		Resume:      about.Title,
		CreatedAt:   about.CreatedAt,
		UpdatedAt:   about.UpdatedAt,
	}
}
