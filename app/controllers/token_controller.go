package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func RenewTokens(c *fiber.Ctx) error {
	// now := time.Now().Unix()

	// claims, err := utils.ExtractTokenMetadata(c)

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"msg": "wtf",
	})
}
