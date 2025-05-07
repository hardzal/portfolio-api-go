package stack

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
)

type StackHandler interface {
	GetStack(c *fiber.Ctx) error
	GetAllStack(c *fiber.Ctx) error
	CreateStack(c *fiber.Ctx) error
	UpdateStack(c *fiber.Ctx) error
	DeleteStack(c *fiber.Ctx) error
}

type StackHandlerImpl struct {
	stackService services.StackService
}

// CreateStack implements StackHandler.
func (s *StackHandlerImpl) CreateStack(c *fiber.Ctx) error {
	var stackDTO models.StackDTO
	if err := utils.ParseBodyAndValidate(c, &stackDTO); err != nil {
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

	url, err := utils.UploadToCloudinary(file, "stacks")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	newStack, err := s.stackService.CreateStack(&stackDTO, &url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create stack",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Sucess created data",
		"data":    newStack,
	})
}

// DeleteStack implements StackHandler.
func (s *StackHandlerImpl) DeleteStack(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	if err := s.stackService.DeleteStack(uint(id)); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  fiber.StatusNoContent,
		"message": "Success delete the stack",
	})
}

// GetAllStack implements StackHandler.
func (s *StackHandlerImpl) GetAllStack(c *fiber.Ctx) error {
	stacks, err := s.stackService.GetAllStack()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get all data",
		"data":    stacks,
	})
}

// GetStack implements StackHandler.
func (s *StackHandlerImpl) GetStack(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	stack, err := s.stackService.GetStack(uint(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "stack not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get the stack",
		"data":    stack,
	})
}

// UpdateStack implements StackHandler.
func (s *StackHandlerImpl) UpdateStack(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))
	var stackDTO models.StackDTO
	if err := utils.ParseBodyAndValidate(c, &stackDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileHeader, err := c.FormFile("image")
	var imageURL string
	if err == nil && fileHeader != nil {
		uploadedURL, err := utils.UploadToCloudinary(fileHeader, "stacks")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "upload failed")
		}
		imageURL = uploadedURL
	}

	updatedProject, err := s.stackService.UpdateStack(uint(id), &stackDTO, imageURL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success updated the stack",
		"data":    updatedProject,
	})
}

func NewStackHandler(stack services.StackService) StackHandler {
	return &StackHandlerImpl{stackService: stack}
}
