package middlewares

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func FileUploadMiddleware(c *fiber.Ctx) error {

	file, err := c.FormFile("image")

	if err != nil {
		log.Println("Error in uploading Image : ", err)

		return c.JSON(fiber.Map{"status": 500, "message": "Server error", "data": nil})

	}
	destination := fmt.Sprintf("./uploads/%s", file.Filename)

	if err := c.SaveFile(file, destination); err != nil {
		return err
	}
	c.Locals("filePath", file.Filename)
	c.Locals("file", file)

	return c.Next()

}
