package routers

import (
	"FoodLine/src/delivery/rest/controllers/auth"
	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(app *fiber.App) {
	app.Post("/auth/register", auth.RegisterUser)
	app.Post("/auth/login", auth.LoginUser)
}
