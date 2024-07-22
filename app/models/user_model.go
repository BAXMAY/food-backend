package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id" bson:"id"`
	CreatedAt    time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" bson:"updated_at"`
	Email        string    `json:"email" bson:"email"`
	PasswordHash string    `json:"password_hash" bson:"password_hash,omitempty"`
	UserStatus   int       `json:"user_status" bson:"user_status"`
	UserRole     string    `json:"user_role" bson:"user_role"`
}
