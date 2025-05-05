package auth

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/models"
	"github.com/hardzal/portfolio-api-go/services"
	"github.com/hardzal/portfolio-api-go/utils"
)

type AuthHandler struct {
	AuthService services.AuthService
}

func NewAuthHandler(authService services.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) LoginHandler(ctx *fiber.Ctx) error {
	var userLoginDTO models.UserLoginDTO

	if err := utils.ParseBodyAndValidate(ctx, &userLoginDTO); err != nil {
		return err
	}

	token, err := h.AuthService.LoginUser(userLoginDTO)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Login failed",
			"error":   errors.New("email or password wrong"),
		})
	}
	currentTime := time.Now()
	newTime := currentTime.Add(time.Hour)

	return ctx.Status(fiber.StatusOK).JSON(&models.AuthResponse{
		Token:     token,
		ExpiresAt: &newTime,
	})
}
