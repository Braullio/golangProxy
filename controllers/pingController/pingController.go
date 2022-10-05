package pingController

import (
	"github.com/gofiber/fiber/v2"
	"golangProxy/models"
)

func Index(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(&models.Ping{Msg: "pong"})
}
