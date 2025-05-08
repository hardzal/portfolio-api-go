package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/about"
	"github.com/hardzal/portfolio-api-go/middlewares"
)

func AboutRoutes(app fiber.Router, handler about.AboutHandler) {
	app.Get("/:id", handler.GetAbout)
	app.Post("/", middlewares.Auth, handler.CreateAbout)
	app.Put("/:id", middlewares.Auth, handler.UpdateAbout)
}
