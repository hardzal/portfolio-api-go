package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/handlers/auth"
)

func AuthRoutes(app fiber.Router, handler *auth.AuthHandler) {
	app.Post("/login", handler.LoginHandler)
}
