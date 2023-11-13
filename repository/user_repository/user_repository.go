package userrepository

import (
	"kosei-lms-api/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		DB: db,
	}
}

func (s *UserRepository) GetByID(id string) error {
	var user models.User
	result := s.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
