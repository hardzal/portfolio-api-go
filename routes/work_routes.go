package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/work"
	"github.com/hardzal/portfolio-api-go/middlewares"
)

func WorkRoutes(app fiber.Router, handler work.WorkHandler) {
	app.Get("/:id", handler.GetWork)
	app.Get("/", handler.GetAllWork)
	app.Post("/", middlewares.Auth, handler.CreateWork)
	app.Put("/:id", middlewares.Auth, handler.UpdateWork)
	app.Delete("/:id", middlewares.Auth, handler.DeleteWork)
}
