package models

import (
	// "strconv"
	// "strings"
	// "time"

	"strconv"
	"strings"
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
	CreatedAt time.Time      `json:"created_at"`              // Automatically managed by GORM for creation time
	UpdatedAt time.Time      `json:"updated_at"`              // Automatically managed by GORM for update time
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"` // Automatically managed by GORM
}
type CustomersResponse struct {
	Customers []Customer `json:"customers"`
	Count     int        `json:"count"`
	Search    string     `json:"search"`
	Order     string     `json:"order"`
	Offset    string     `json:"offset"`
	Limit     string     `json:"limit"`
}
type CustomerResponse struct {
	Customer Customer `json:"customer"`
}

func GetCustomers(db *gorm.DB, search string, order string, offset string, limit string) ([]Customer, error) {
	var customers []Customer
	var orderBy = strings.Replace(order, "-", " ", -1)
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		return nil, err
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		return nil, err
	}
	result := db.Where("CONCAT(name, ' ', gender, ' ', email, ' ', phone, ' ', address) LIKE ?", "%"+search+"%").Order(orderBy).Limit(limitInt).Offset(offsetInt).Find(&customers)
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
