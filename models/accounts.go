package models

import (
	// "strconv"
	// "strings"
	// "time"

	"time"

	"gorm.io/gorm"
)

// !!!be careful, CreatedAt,UpdatedAt,DeletedAt are automatically managed by GORM
// when you delete, DeletedAt will be not null and set to time.Now(), you auto can't select any more
type Account struct {
	ID        uint64         `json:"id" gorm:"primary_key;auto_increment"`
	Username  string         `json:"username" gorm:"unique;not null"`
	Password  string         `json:"password" gorm:"not null"`
	Token     string         `json:"token" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time      `json:"updated_at"` // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at"` // Automatically managed by GORM
}

type AccountResponse struct {
	ID        uint64         `json:"id"`
	Username  string         `json:"username"`
	Token     string         `json:"token"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

type AccountChangePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type AccountLoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetAccounts(db *gorm.DB) ([]Account, error) {
	var Accounts []Account
	err := db.Find(&Accounts).Error
	if err != nil {
		return nil, err
	}
	return Accounts, nil
}

func GetAccountById(db *gorm.DB, id uint64) (Account, error) {
	var account Account
	err := db.First(&account, id).Error
	if err != nil {
		return account, err
	}
	return account, nil
}

func GetAccountByUsername(db *gorm.DB, username string) (Account, error) {
	var account Account
	err := db.First(&account, "username = ?", username).Error
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

func DeleteAccount(db *gorm.DB, id uint64) error {
	//two way for delete are same
	// err := db.Delete(&Account{}, id).Error
	err := db.Where("id = ?", id).Delete(&Account{}).Error
	if err != nil {
		return err
	}
	return nil
}
