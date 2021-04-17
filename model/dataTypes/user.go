package dataTypes

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	UserId uuid.UUID `json:"-"`
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	Username string `json:"username"`
	DateCreated time.Time `json:"date_created"`
}
