package auth_controller

import (
	authrepository "kosei-lms-api/repository/auth_repository"

	"gorm.io/gorm"
)

type AuthController struct {
	DB   *gorm.DB
	Auth authrepository.Authrepository
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{
		DB:   DB,
		Auth: authrepository.NewAuthRepository(DB),
	}
}
