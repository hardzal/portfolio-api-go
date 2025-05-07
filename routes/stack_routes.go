package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/stack"
	"github.com/hardzal/portfolio-api-go/middlewares"
)

func StackRoutes(app fiber.Router, handler stack.StackHandler) {
	app.Get("/:id", handler.GetStack)
	app.Get("/", handler.GetAllStack)
	app.Post("/", middlewares.Auth, handler.CreateStack)
	app.Put("/:id", middlewares.Auth, handler.UpdateStack)
	app.Delete("/:id", middlewares.Auth, handler.DeleteStack)
}
