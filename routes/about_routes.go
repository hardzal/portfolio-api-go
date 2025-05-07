package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/about"
)

func AboutRoutes(app fiber.Router, handler about.AboutHandler) {
	app.Get("/:id", handler.GetAbout)
	app.Get("/", handler.GetAllAbout)
	app.Post("/", handler.CreateAbout)
	app.Put("/:id", handler.UpdateAbout)
	app.Delete("/:id", handler.DeleteAbout)
}
