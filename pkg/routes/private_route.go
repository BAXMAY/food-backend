package routes

import (
	"github.com/BAXMAY/food-backend/app/controllers"
	"github.com/BAXMAY/food-backend/pkg/middleware"
	"github.com/gofiber/fiber/v2"
)

func PrivateRoute(a *fiber.App) {
	route := a.Group("/api/v1")

	route.Post("/book", middleware.JWTProtected(), controllers.CreateBook)
	route.Post("/user/sign/out", middleware.JWTProtected(), controllers.UserSignOut)
	route.Post("/token/renew", middleware.JWTProtected(), controllers.RenewTokens)

	route.Put("/book", middleware.JWTProtected(), controllers.UpdateBook)

	route.Delete("/book", middleware.JWTProtected(), controllers.DeleteBook)
}
