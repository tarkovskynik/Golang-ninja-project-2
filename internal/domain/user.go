package domain

import (
	"time"
)

type User struct {
	// Email is unique value in DB
	Email     string    `json:"email,omitempty" bson:"email" example:"bg@example.com"`
	Password  string    `json:"password,omitempty" bson:"password" example:"Bill"`
	Token     string    `json:"-" bson:"token"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
}
