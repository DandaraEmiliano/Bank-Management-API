package services

import (
	"bank-management/models"
	"errors"

	"gorm.io/gorm"
)

func Deposit(db *gorm.DB, accountID uint, amount float64) error {
	var account models.Account
	if err := db.First(&account, accountID).Error; err != nil {
		return errors.New("account not found")
	}

	account.Balance += amount
	db.Save(&account)
	return nil
}

func Withdraw(db *gorm.DB, accountID uint, amount float64) error {
	var account models.Account
	if err := db.First(&account, accountID).Error; err != nil {
		return errors.New("account not found")
	}

	if account.Balance < amount {
		return errors.New("insufficient balance")
	}

	account.Balance -= amount
	db.Save(&account)
	return nil
}
