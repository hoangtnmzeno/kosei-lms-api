package models

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/google/uuid"
)

// role == 0 || 1 : user ? admin

// type == 0 || 1 : user ? teacher
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name      string    `gorm:"uniqueIndex;not null"`
	Email     string    `gorm:"not null"`
	Password  string    `gorm:"not null"`
	Role      int8      `gorm:"not null"`
	Photo     string    `gorm:"not null"`
	Type      int8      `gorm:"not null"`
	Delete    int8      `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type SignUpInput struct {
	Name            string                `form:"name"`
	Email           string                `form:"email"`
	Password        string                `form:"password"`
	PasswordConfirm string                `form:"passwordConfirm"`
	Photo           *multipart.FileHeader `form:"photo"`
}

type SignInInput struct {
	Name     string `json:"name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
	Otp      string `json:"otp"`
}

type UserResponse struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Role      int8      `json:"role,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	Type      int8      `json:"provider"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type SignPhoto struct {
	Photo *multipart.FileHeader `form:"photo"`
}

func ValidateUser(u SignUpInput) error {
	if u.Name == "" {
		return errors.New("name invalid!")
	}

	if u.Email == "" {
		return errors.New("email invalid!")
	}

	if u.Password == "" {
		return errors.New("password invalid!")
	}

	if len(u.Password) < 8 {
		return errors.New("password min 8 character!")
	}

	if u.PasswordConfirm == "" {
		return errors.New("name invalid!")
	}

	return nil
}
