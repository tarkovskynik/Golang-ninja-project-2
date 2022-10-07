package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	// Email is unique value in DB
	Email     string    `json:"email,omitempty" bson:"email" example:"bg@example.com"`
	Password  string    `json:"password,omitempty" bson:"password" example:"Bill"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
}
