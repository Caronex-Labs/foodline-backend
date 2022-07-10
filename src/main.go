package main

import (
	"FoodLine/src/delivery/rest/routers"
	"FoodLine/src/utils"
	"FoodLine/src/utils/database"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initializations
	config := utils.NewConfig()
	mongoClient = database.NewMongoClient(config)

	app := fiber.New()

	routers.AuthRoutes(app)

	err := app.Listen(":6000")
	if err != nil {
		return
	}
}
