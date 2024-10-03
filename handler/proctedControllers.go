package handler

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func Welcome(c *fiber.Ctx) error {
	x := 10
	return c.JSON(fiber.Map{
		"status":  200,
		"message": fmt.Sprintf("welcome, %v", x),
	})
}
