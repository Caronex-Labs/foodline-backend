package models

import (
	"time"
)

type User struct {
	ID                string `json:"id,omitempty" bson:"_id,omitempty"`
	FirstName         string `json:"first_name" validate:"required"`
	LastName          string `json:"last_name" validate:"required"`
	Location          string `json:"location" validate:"required"`
	Email             string `json:"email" validate:"required,email"`
	Password          string `json:"password" validate:"required"`
	ProfilePictureURL string `json:"profile_picture_url,omitempty" validate:"url"`

	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
