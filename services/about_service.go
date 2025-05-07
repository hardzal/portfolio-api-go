package services

import (
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/repositories"
)

type AboutService interface {
	GetAbout(id uint) (*models.AboutResponse, error)
	CreateAbout(about *models.AboutDTO, image *string) (*models.AboutResponse, error)
	UpdateAbout(id uint, about *models.AboutDTO, image *string) (*models.AboutResponse, error)
}

type aboutService struct {
	aboutRepo repositories.AboutRepository
}

func NewAboutService(repo repositories.AboutRepository) AboutService {
	return &aboutService{aboutRepo: repo}
}

// CreateAbout implements AboutService.
func (a *aboutService) CreateAbout(about *models.AboutDTO, image *string) (*models.AboutResponse, error) {
	newAbout := &models.About{
		Title:       about.Title,
		Profession:  about.Profession,
		Description: about.Description,
		Location:    about.Location,
		IsAvailable: about.IsAvailable,
		Handphone:   about.Handphone,
		Email:       about.Email,
		Resume:      &about.Resume,
	}
	createdAbout, err := a.aboutRepo.CreateAbout(newAbout)
	if err != nil {
		return nil, err
	}

	return a.mapToResponse(createdAbout), nil
}

// GetAbout implements AboutService.
func (a *aboutService) GetAbout(id uint) (*models.AboutResponse, error) {
	about, err := a.aboutRepo.GetAbout(id)
	if err != nil {
		return nil, err
	}

	return a.mapToResponse(about), nil
}

// UpdateAbout implements AboutService.
func (a *aboutService) UpdateAbout(id uint, about *models.AboutDTO, image *string) (*models.AboutResponse, error) {
	dataAbout, err := a.aboutRepo.GetAbout(id)
	if err != nil {
		return nil, err
	}

	updatedAbout := &models.About{
		ID:          dataAbout.ID,
		Title:       about.Title,
		Profession:  about.Profession,
		Description: about.Description,
		Location:    about.Location,
		IsAvailable: about.IsAvailable,
		Handphone:   about.Handphone,
		Email:       about.Email,
		Resume:      &about.Resume,
	}

	newUpdateAbout, err := a.aboutRepo.UpdateAbout(updatedAbout)

	if err != nil {
		return nil, err
	}

	return a.mapToResponse(newUpdateAbout), nil
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
		Resume:      *about.Resume,
		CreatedAt:   about.CreatedAt,
		UpdatedAt:   about.UpdatedAt,
	}
}
