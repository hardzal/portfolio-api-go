package project

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
)

type ProjectHandler interface {
	GetProject(c *fiber.Ctx) error
	GetAllProject(c *fiber.Ctx) error
	CreateProject(c *fiber.Ctx) error
	UpdateProject(c *fiber.Ctx) error
	DeleteProject(c *fiber.Ctx) error
}
type projectHandler struct {
	projectService services.ProjectService
}

// CreateProject implements ProjectHandler.
func (p *projectHandler) CreateProject(c *fiber.Ctx) error {
	var projectDTO models.ProjectDTO
	if err := utils.ParseBodyAndValidate(c, &projectDTO); err != nil {
		return err
	}

	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Image file required",
			"error":   err,
		})
	}

	url, err := utils.UploadToCloudinary(file, "projects")
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	newProject, err := p.projectService.CreateProject(&projectDTO, &url)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create project",
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"status":  fiber.StatusCreated,
		"message": "Sucess created data",
		"data":    newProject,
	})
}

// DeleteProject implements ProjectHandler.
func (p *projectHandler) DeleteProject(c *fiber.Ctx) error {
	id := c.Params("id")
	if err := p.projectService.DeleteProject(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusNoContent).JSON(fiber.Map{
		"status":  fiber.StatusNoContent,
		"message": "Success delete the project",
	})
}

// GetAllProject implements ProjectHandler.
func (p *projectHandler) GetAllProject(c *fiber.Ctx) error {
	projects, err := p.projectService.GetAllProjects()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get all data",
		"data":    projects,
	})
}

// GetProject implements ProjectHandler.
func (p *projectHandler) GetProject(c *fiber.Ctx) error {
	id := c.Params("id")
	project, err := p.projectService.GetProject(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "project not found",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success get the project",
		"data":    project,
	})
}

// UpdateProject implements ProjectHandler.
func (p *projectHandler) UpdateProject(c *fiber.Ctx) error {
	id := c.Params("id")
	var projectDTO models.ProjectDTO
	if err := utils.ParseBodyAndValidate(c, &projectDTO); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	fileHeader, err := c.FormFile("image")
	var imageURL string
	if err == nil && fileHeader != nil {
		uploadedURL, err := utils.UploadToCloudinary(fileHeader, "projects")
		if err != nil {
			return fiber.NewError(fiber.StatusInternalServerError, "upload failed")
		}
		imageURL = uploadedURL
	}

	updatedProject, err := p.projectService.UpdateProject(id, &projectDTO, &imageURL)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  fiber.StatusOK,
		"message": "success updated the project",
		"data":    updatedProject,
	})
}

func NewProjectHandler(service services.ProjectService) ProjectHandler {
	return &projectHandler{projectService: service}
}
