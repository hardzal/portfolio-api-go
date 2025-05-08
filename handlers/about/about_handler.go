package about

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
)

type AboutHandler interface {
	GetAbout(c *fiber.Ctx) error
	CreateAbout(c *fiber.Ctx) error
	UpdateAbout(c *fiber.Ctx) error
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
			"error": "failed to create about",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Sucess created data",
		"data":    newAbout,
	})
}

// GetAbout implements AboutHandler.
func (a *AboutHandlerImpl) GetAbout(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	about, err := a.AboutService.GetAbout(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "about not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get the about",
		"data":    about,
	})
}

// UpdateAbout implements AboutHandler.
func (a *AboutHandlerImpl) UpdateAbout(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var aboutDTO models.AboutDTO
	if err := utils.ParseBodyAndValidate(c, &aboutDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileHeader, err := c.FormFile("image")
	var imageURL string
	if err == nil && fileHeader != nil {
		uploadedURL, err := utils.UploadToCloudinary(fileHeader, "about")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "upload failed")
		}
		imageURL = uploadedURL
	}

	updatedAbout, err := a.AboutService.UpdateAbout(uint(id), &aboutDTO, &imageURL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success updated the project",
		"data":    updatedAbout,
	})
}

func NewAboutHandler(aboutService services.AboutService) AboutHandler {
	return &AboutHandlerImpl{AboutService: aboutService}
}
