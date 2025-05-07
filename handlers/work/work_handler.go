package work

import (
	"github.com/gofiber/fiber/v2"
	"github.com/hardzal/portfolio-api-go/services"
)

type WorkHandler interface {
	GetWork(c *fiber.Ctx) error
	GetAllWork(c *fiber.Ctx) error
	CreateWork(c *fiber.Ctx) error
	UpdateWork(c *fiber.Ctx) error
	DeleteWork(c *fiber.Ctx) error
}

type WorkHandlerImpl struct {
	workService services.WorkService
}

// CreateWork implements WorkHandler.
func (w *WorkHandlerImpl) CreateWork(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteWork implements WorkHandler.
func (w *WorkHandlerImpl) DeleteWork(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetAllWork implements WorkHandler.
func (w *WorkHandlerImpl) GetAllWork(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetWork implements WorkHandler.
func (w *WorkHandlerImpl) GetWork(c *fiber.Ctx) error {
	panic("unimplemented")
}

// UpdateWork implements WorkHandler.
func (w *WorkHandlerImpl) UpdateWork(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewWorkHandler(work services.WorkService) WorkHandler {
	return &WorkHandlerImpl{workService: work}
}
