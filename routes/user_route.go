package routes

import (
	"FoodLine/controllers" //add this

	"github.com/gofiber/fiber/v2"
)

func UserRoute(app *fiber.App) {
	app.Post("/auth/register", controllers.RegisterUser)
	app.Post("/auth/login", controllers.LoginUser)
}
