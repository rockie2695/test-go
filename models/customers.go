package models

import (
	// "strconv"
	// "strings"
	// "time"

	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID        uint64         `json:"id" gorm:"primary_key;auto_increment"`
	Name      string         `json:"name" gorm:"not null"`
	Gender    string         `json:"gender" gorm:"not null"`
	Email     string         `json:"email" gorm:"unique;not null"`
	Phone     string         `json:"phone" gorm:"not null"`
	Address   string         `json:"address" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"` // Automatically managed by GORM for creation time
	UpdatedAt time.Time      `json:"updated_at"` // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at"` // Automatically managed by GORM
}
type CustomersResponse struct {
	Customers []Customer `json:"customers"`
	Count     int        `json:"count"`
}
type CustomerResponse struct {
	Customer Customer `json:"customer"`
}

func GetCustomers(db *gorm.DB) ([]Customer, error) {
	var customers []Customer
	result := db.Find(&customers)
	if result.Error != nil {
		return nil, result.Error
	}
	return customers, nil
}

func GetCustomerById(db *gorm.DB, id uint64) (Customer, error) {
	var customer Customer
	result := db.First(&customer, id)
	if result.Error != nil {
		return Customer{}, result.Error
	}
	return customer, nil
}

func CreateCustomer(db *gorm.DB, customer Customer) (Customer, error) {
	result := db.Create(&customer)
	if result.Error != nil {
		return Customer{}, result.Error
	}
	return customer, nil
}

func UpdateCustomer(db *gorm.DB, id uint64, customer Customer) (Customer, error) {
	result := db.Model(&customer).Where("id = ?", id).Updates(&customer)
	if result.Error != nil {
		return Customer{}, result.Error
	}
	return customer, nil
}

func DeleteCustomer(db *gorm.DB, id uint64) error {
	result := db.Delete(&Customer{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
