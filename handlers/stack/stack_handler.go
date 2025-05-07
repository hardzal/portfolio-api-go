package stack

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/services"
)

type StackHandler interface {
	GetStack(c *fiber.Ctx) error
	GetAllStack(c *fiber.Ctx) error
	CreateStack(c *fiber.Ctx) error
	UpdateStack(c *fiber.Ctx) error
	DeleteStack(c *fiber.Ctx) error
}

type StackHandlerImpl struct {
	stackService services.StackService
}

// CreateStack implements StackHandler.
func (s *StackHandlerImpl) CreateStack(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteStack implements StackHandler.
func (s *StackHandlerImpl) DeleteStack(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllStack implements StackHandler.
func (s *StackHandlerImpl) GetAllStack(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetStack implements StackHandler.
func (s *StackHandlerImpl) GetStack(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateStack implements StackHandler.
func (s *StackHandlerImpl) UpdateStack(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewStackHandler(stack services.StackService) StackHandler {
	return &StackHandlerImpl{stackService: stack}
}
