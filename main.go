package main

import (
	"FoodLine/configs"
	"FoodLine/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	configs.ConnectDB()

	routes.UserRoute(app) //add this

	app.Listen(":6000")
}
