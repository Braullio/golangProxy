package transparencyPortalController

import "github.com/gofiber/fiber/v2"

func Show(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusOK)
}
