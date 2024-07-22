package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func UserSignUp(c *fiber.Ctx) error {
	// signUp := &models.SignUp{}

	// return c.JSON(fiber.Map{
	// 	"error": false,
	// 	"msg":   nil,
	// 	"user":  user,
	// })

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"msg": "wtf",
	})
}

func UserSignIn(c *fiber.Ctx) error {

	// // Generate a new pair of access and refresh tokens.
	// tokens, err := utils.GenerateNewTokens(foundedUser.ID.String(), credentials)
	// if err != nil {
	// 	// Return status 500 and token generation error.
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// return c.JSON(fiber.Map{
	// 	"error": false,
	// 	"msg":   nil,
	// 	"tokens": fiber.Map{
	// 		"access":  tokens.Access,
	// 		"refresh": tokens.Refresh,
	// 	},
	// })

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"msg": "wtf",
	})
}

func UserSignOut(c *fiber.Ctx) error {
	// claims, err := utils.ExtractTokenMetadata(c)

	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
	// 		"error": true,
	// 		"msg":   err.Error(),
	// 	})
	// }

	// userID := claims.UserID.String()

	// return c.SendStatus(fiber.StatusNoContent)

	return c.Status(fiber.StatusNotImplemented).JSON(fiber.Map{
		"msg": "wtf",
	})
}
