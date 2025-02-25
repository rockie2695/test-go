package models

import (
	// "strconv"
	// "strings"
	// "time"

	"time"

	"gorm.io/gorm"
)

type Account struct {
	ID        int64          `json:"id" gorm:"primary_key;auto_increment"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time      `json:"updated_at"` // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at"` // Automatically managed by GORM
}

func GetAccounts(db *gorm.DB) ([]Account, error) {
	var Accounts []Account
	err := db.Find(&Accounts).Error
	if err != nil {
		return nil, err
	}
	return Accounts, nil
}

func GetAccountById(db *gorm.DB, id int64) (Account, error) {
	var account Account
	err := db.First(&account, id).Error
	if err != nil {
		return account, err
	}
	return account, nil
}

func CreateAccount(db *gorm.DB, account Account) (Account, error) {
	err := db.Create(&account).Error
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func UpdateAccount(db *gorm.DB, account Account) (Account, error) {
	//default save is base on id
	err := db.Save(&account).Error
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

func DeleteAccount(db *gorm.DB, id int64) error {
	//two way for delete are same
	// err := db.Delete(&Account{}, id).Error
	err := db.Where("id = ?", id).Delete(&Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
