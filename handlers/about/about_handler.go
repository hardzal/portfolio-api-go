package about

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/services"
)

type AboutHandler interface {
	GetAbout(c *fiber.Ctx) error
	GetAllAbout(c *fiber.Ctx) error
	CreateAbout(c *fiber.Ctx) error
	UpdateAbout(c *fiber.Ctx) error
	DeleteAbout(c *fiber.Ctx) error
}

type AboutHandlerImpl struct {
	AboutService services.AboutService
}

// CreateAbout implements AboutHandler.
func (a *AboutHandlerImpl) CreateAbout(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteAbout implements AboutHandler.
func (a *AboutHandlerImpl) DeleteAbout(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAbout implements AboutHandler.
func (a *AboutHandlerImpl) GetAbout(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllAbout implements AboutHandler.
func (a *AboutHandlerImpl) GetAllAbout(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateAbout implements AboutHandler.
func (a *AboutHandlerImpl) UpdateAbout(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewAboutHandler(aboutService services.AboutService) AboutHandler {
	return &AboutHandlerImpl{AboutService: aboutService}
}
