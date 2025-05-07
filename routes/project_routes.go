package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/project"
)

func ProjectRoutes(app fiber.Router, handler project.ProjectHandler) {
	app.Get("/:id", handler.GetProject)
	app.Get("/", handler.GetAllProject)
	app.Post("/", handler.CreateProject)
	app.Put("/:id", handler.UpdateProject)
	app.Delete("/:id", handler.DeleteProject)
}
