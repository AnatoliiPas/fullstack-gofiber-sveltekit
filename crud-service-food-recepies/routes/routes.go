package routes

import (
	"crud-service-food-recepies/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
	app.Get("food-recepies/api/v1/index", controllers.Index)
}
