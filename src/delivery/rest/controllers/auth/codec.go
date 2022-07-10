package auth

import (
	"FoodLine/src/internal/models"
	"github.com/gofiber/fiber/v2"
)

type RegisterUserRequest struct {
	FirstName         string `json:"first_name" validate:"required"`
	LastName          string `json:"last_name" validate:"required"`
	Location          string `json:"location" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	ProfilePictureURL string `json:"profile_picture_url,omitempty" validate:"url"`
}

type RegisterUserResponse struct {
	User *models.User
}

type LoginUserRequest struct {
}

type LoginUserResponse struct {
}

type ErrorResponse struct {
	Message string     `json:"message"`
	Data    *fiber.Map `json:"data"`
}
