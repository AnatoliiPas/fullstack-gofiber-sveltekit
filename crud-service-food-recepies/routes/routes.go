package routes

import (
	"crud-service-food-recepies/controllers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
	app.Get("food-recepies/api/v1/recipes/index", controllers.Index)
	app.Get("food-recepies/api/v1/recipes/about", controllers.About)
	app.Get("food-recepies/api/v1/recipes", controllers.GetAll)
	app.Get("food-recepies/api/v1/recipes/:id", controllers.GetRecipe)
	app.Post("food-recepies/api/v1/recipes", controllers.CreateRecipe)
	app.Patch("food-recepies/api/v1/recipes/:id", controllers.UpdateRecipe)
	app.Delete("food-recepies/api/v1/recipes/:id", controllers.DeleteRecipe)
}
