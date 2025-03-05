package controllers

import (
	"errors"
	// "fmt"
	"net/http"
	"strconv"
	"test-go/database"
	// "test-go/middleware"
	"test-go/models"
	// "time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CustomersAutoMigrate() {
	database.Db.AutoMigrate(&models.Customer{})
}

// @Security     BearerAuth
// @Summary      Get all customers
// @Description  Get all customers
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        search   query      string  false  "Search"
// @Param        order    query      string  false  "Order"
// @Param        offset   query      string  false  "Offset"
// @Param        limit    query      string  false  "Limit"
// @Success      200  {array}   models.CustomersResponse
// @Failure      500  {object}  models.HTTPError
// @Router       /customers [get]
func GetCustomers(c *gin.Context) {
	var search = c.DefaultQuery("search", "%")
	var order = c.DefaultQuery("order", "id-desc")
	var offset = c.DefaultQuery("offset", "0")
	var limit = c.DefaultQuery("limit", "20")
	customers, err := models.GetCustomers(database.Db, search, order, offset, limit)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customers": customers, "count": len(customers), "search": search, "order": order, "offset": offset, "limit": limit})
}

// @Security     BearerAuth
// @Summary      Get customer by id
// @Description  Get customer by id
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "customer id"
// @Success      200  {object}   models.CustomerResponse
// @Failure      500  {object}  models.HTTPError
// @Router       /customers/{id} [get]
func GetCustomerById(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	customer, err := models.GetCustomerById(database.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// @Security     BearerAuth
// @Summary      Create customer
// @Description  Create customer
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        customer   body      models.Customer  true  "Customer"
// @Success      201  {object}   models.CustomerResponse
// @Failure      500  {object}  models.HTTPError
// @Router       /customers [post]
func CreateCustomer(c *gin.Context) {
	var customer models.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	customer, err := models.CreateCustomer(database.Db, customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// @Security     BearerAuth
// @Summary      Update customer
// @Description  Update customer
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "customer id"
// @Param        customer   body      models.Customer  true  "Customer"
// @Success      200  {object}   models.CustomerResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /customers/{id} [put]
func UpdateCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	customer, err := models.GetCustomerById(database.Db, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "customer not found"})
			return
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	customer, err = models.UpdateCustomer(database.Db, customer.ID, customer)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"customer": customer})
}

// @Security     BearerAuth
// @Summary      Delete customer
// @Description  Delete customer
// @Tags         customers
// @Accept       json
// @Produce      json
// @Param        id   path      uint64  true  "customer id"
// @Success      200  {object}  models.HTTPResponse
// @Failure      400  {object}  models.HTTPError
// @Failure      500  {object}  models.HTTPError
// @Router       /customers/{id} [delete]
func DeleteCustomer(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	err = models.DeleteCustomer(database.Db, id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
}
