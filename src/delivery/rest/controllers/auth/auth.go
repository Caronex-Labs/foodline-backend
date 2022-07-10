package auth

import (
	"FoodLine/src/internal/usecase"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

func RegisterUser(c *fiber.Ctx) error {
	// de-serialize
	var req RegisterUserRequest
	//validate the request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(http.StatusBadRequest).JSON(ErrorResponse{Message: "error", Data: &fiber.Map{"data": err.Error()}})
	}

	// call usecase
	authUseCase := usecase.NewAuthUseCase()
	user, err := authUseCase.RegisterUser(c, req.FirstName, req.LastName, req.Email, req.Password, req.Location, req.ProfilePictureURL)
	if err != nil {
		return err
	}

	// serialize
	resp := &RegisterUserResponse{
		User: user,
	}

	// respond
	return c.Status(http.StatusCreated).JSON(resp)
}

func LoginUser(c *fiber.Ctx) error {
	// de-serialize
	// call usecase
	// serialize
	// respond
}
