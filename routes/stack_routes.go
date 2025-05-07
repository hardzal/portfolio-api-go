package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/stack"
)

func StackRoutes(app fiber.Router, handler stack.StackHandler) {
	app.Get("/:id", handler.GetStack)
	app.Get("/", handler.GetAllStack)
	app.Post("/", handler.CreateStack)
	app.Put("/:id", handler.UpdateStack)
	app.Delete("/:id", handler.DeleteStack)
}
