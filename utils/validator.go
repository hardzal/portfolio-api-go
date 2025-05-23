package utils

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

var validate = validator.New()

func Validate(payload interface{}) *fiber.Error {
	err := validate.Struct(payload)
	if err != nil {
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("`%v` with value `%v` doesnt satify the `%v` constraint", err.Field(), err.Value(), err.Tag()))
		}

		return &fiber.Error{Code: fiber.StatusBadRequest, Message: strings.Join(errors, ",")}
	}

	return nil
}
