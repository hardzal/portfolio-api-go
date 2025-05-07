package work

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
)

type WorkHandler interface {
	GetWork(c *fiber.Ctx) error
	GetAllWork(c *fiber.Ctx) error
	CreateWork(c *fiber.Ctx) error
	UpdateWork(c *fiber.Ctx) error
	DeleteWork(c *fiber.Ctx) error
}

type WorkHandlerImpl struct {
	workService services.WorkService
}

// CreateWork implements WorkHandler.
func (w *WorkHandlerImpl) CreateWork(c *fiber.Ctx) error {
	var workDTO models.WorkDTO
	if err := utils.ParseBodyAndValidate(c, &workDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image file required",
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

	url, err := utils.UploadToCloudinary(file, "works")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	newWork, err := w.workService.CreateWork(&workDTO, &url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create work",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Sucess created data",
		"data":    newWork,
	})
}

// DeleteWork implements WorkHandler.
func (w *WorkHandlerImpl) DeleteWork(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := w.workService.DeleteWork(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  fiber.StatusNoContent,
		"message": "Success delete the work",
	})
}

// GetAllWork implements WorkHandler.
func (w *WorkHandlerImpl) GetAllWork(c *fiber.Ctx) error {
	works, err := w.workService.GetAllWork()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get all data",
		"data":    works,
	})
}

// GetWork implements WorkHandler.
func (w *WorkHandlerImpl) GetWork(c *fiber.Ctx) error {
	id := c.Params("id")
	work, err := w.workService.GetWork(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "work not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get the work",
		"data":    work,
	})
}

// UpdateWork implements WorkHandler.
func (w *WorkHandlerImpl) UpdateWork(c *fiber.Ctx) error {
	id := c.Params("id")
	var workDTO models.WorkDTO
	if err := utils.ParseBodyAndValidate(c, &workDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileHeader, err := c.FormFile("image")
	var imageURL string
	if err == nil && fileHeader != nil {
		uploadedURL, err := utils.UploadToCloudinary(fileHeader, "works")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "upload failed")
		}
		imageURL = uploadedURL
	}

	updatedWork, err := w.workService.UpdateWork(id, &workDTO, &imageURL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success updated the work",
		"data":    updatedWork,
	})
}

func NewWorkHandler(work services.WorkService) WorkHandler {
	return &WorkHandlerImpl{workService: work}
}
