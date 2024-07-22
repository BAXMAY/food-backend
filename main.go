package main

import (
	"os"

	"github.com/BAXMAY/food-backend/pkg/configs"
	"github.com/BAXMAY/food-backend/pkg/middleware"
	"github.com/BAXMAY/food-backend/pkg/routes"
	"github.com/BAXMAY/food-backend/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	config := configs.FiberConfig()

	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.PublicRoute(app)
	routes.PrivateRoute(app)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
