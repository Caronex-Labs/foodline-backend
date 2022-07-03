package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id                primitive.ObjectID `json:"id,omitempty" validate:"unique"`
	FirstName         string             `json:"first_name" validate:"required"`
	LastName          string             `json:"last_name" validate:"required"`
	Location          string             `json:"location" validate:"required"`
	Email             string             `json:"email" validate:"required,email,unique"`
	Password          string             `json:"password" validate:"required"`
	ProfilePictureURL string             `json:"profile_picture_url,omitempty" validate:"url"`
	Restraunts        []Restraunt
}

type Restraunt struct {
	Id primitive.ObjectID `json:"id,omitempty" validate:"unique"`
}