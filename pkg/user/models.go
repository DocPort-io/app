package user

import (
	"time"
)

type User struct {
	ID            int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Name          string
	Email         string
	EmailVerified bool
}

type CreateUserRequest struct {
	Name          string
	Email         string
	EmailVerified bool
}
