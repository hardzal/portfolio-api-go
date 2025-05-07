package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/project"
	"github.com/hardzal/portfolio-api-go/middlewares"
)

func ProjectRoutes(app fiber.Router, handler project.ProjectHandler) {
	app.Get("/:id", handler.GetProject)
	app.Get("/", handler.GetAllProject)
	app.Post("/", middlewares.Auth, handler.CreateProject)
	app.Put("/:id", middlewares.Auth, handler.UpdateProject)
	app.Delete("/:id", middlewares.Auth, handler.DeleteProject)
}
