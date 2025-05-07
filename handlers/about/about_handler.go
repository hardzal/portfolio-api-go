package about

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
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
	var aboutDTO models.AboutDTO
	if err := utils.ParseBodyAndValidate(c, &aboutDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Message,
			"error":   err.Error(),
		})
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image file required",
			"error":   err.Error(),
		})
	}

	url, err := utils.UploadToCloudinary(file, "about")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	newAbout, err := a.AboutService.CreateAbout(&aboutDTO, &url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create project",
		})
	}

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
