package routes

import (
	"github.com/BAXMAY/food-backend/app/controllers"
	"github.com/gofiber/fiber/v2"
)

func PublicRoute(a *fiber.App) {
	route := a.Group("api/v1")

	route.Get("/books", controllers.GetBooks)
	route.Get("/book/:id", controllers.GetBook)

	route.Get("/user/sign/up", controllers.UserSignUp)
	route.Get("/user/sign/in", controllers.UserSignIn)
}
