package usecase

import (
	"FoodLine/src/internal/models"
	"github.com/gofiber/fiber/v2"
	"os/user"
)

type authUseCase struct {
}

func (au *authUseCase) RegisterUser(ctx *fiber.Ctx, firstName string, lastName string, email string, password string, location string, profilePictureURL string) (*models.User, error) {
	// create a new user
	// Add created and updated dates
	// Save to db
	// return user
}

func (au *authUseCase) LoginUser(ctx *fiber.Ctx, email string, password string) (*user.User, error) {
	// find if user with that email exists
	// find if the given password matches with that user
	// return user
}

func NewAuthUseCase() *authUseCase {
	return &authUseCase{}
}
