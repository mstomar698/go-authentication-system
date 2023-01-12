package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// type User struct {
//     Id       primitive.ObjectID `json:"id,omitempty"`
//     Name     string             `json:"name,omitempty" validate:"required"`
//     Location string             `json:"location,omitempty" validate:"required"`
//     Title    string             `json:"title,omitempty" validate:"required"`
// }

// Schema for authentication
type User struct {
	ID            primitive.ObjectID `json:"_id,omitempty"`
	First_name    *string            `json:"first_name" validate:"required,min=2,max=100"`
	Last_name     *string            `json:"last_name" validate:"required,min=2,max=100"`
	Password      *string            `json:"Password" validate:"required,min=6"`
	Email         *string            `json:"email" validate:"email,required"`
	Phone         *string            `json:"phone" validate:"required,min=8,max=12"`
	Token         *string            `json:"token"`
	User_type     *string            `json:"user_type" validate:"required,eq=ADMIN|eq=USER"`
	Refresh_token *string            `json:"refresh_token"`
	Created_at    time.Time          `json:"created_at"`
	Updated_at    time.Time          `json:"updated_at"`
	User_id       string             `json:"user_id"`
	Name          string             `json:"name,omitempty"`
	Location      string             `json:"location,omitempty"`
	Title         string             `json:"title,omitempty"`
}
