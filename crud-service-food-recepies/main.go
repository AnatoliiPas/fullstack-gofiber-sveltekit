package main

import (
	"crud-service-food-recepies/database"
	"crud-service-food-recepies/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func init() {
	database.InitDb()
}

func main() {
	app := fiber.New()

	app.Static("/", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:8080, http://127.0.0.1:5173, http://127.0.0.1:5173/*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	routes.GetRoutes(app)

	app.Listen(":8000")
}
