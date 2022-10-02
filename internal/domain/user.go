package domain

import (
	"time"
)

type User struct {
	Email     string    `json:"email,omitempty" bson:"email" example:"bg@example.com"` // RFC 5322 address
	Password  string    `json:"password,omitempty" bson:"password" example:"Bill"`
	Token     string    `json:"-" bson:"token"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
}
