package services

import (
	"bank-management/models"
	"errors"

	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB, user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return errors.New("error creating user")
	}
	return nil
}
