package usercontroller

import (
	userrepository "kosei-lms-api/repository/user_repository"

	"gorm.io/gorm"
)

type UserController struct {
	DB   *gorm.DB
	User userrepository.UserRepository
}

func NewUserController(DB *gorm.DB) UserController {
	return UserController{
		DB:   DB,
		User: userrepository.NewUserRepository(DB),
	}
}
