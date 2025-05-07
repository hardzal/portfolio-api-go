package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/work"
)

func WorkRoutes(app fiber.Router, handler work.WorkHandler) {
	app.Get("/:id", handler.GetWork)
	app.Get("/", handler.GetAllWork)
	app.Post("/", handler.CreateWork)
	app.Put("/:id", handler.UpdateWork)
	app.Delete("/:id", handler.DeleteWork)
}
