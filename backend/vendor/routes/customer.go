package routes

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"net/http"

	"models"
	"utils"
)


var customerTitle = "Customer"

func CreateCustomer(c *gin.Context) {
	var w models.Customer
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetCustomerByPhone(w.Phone)
		if existItem != nil {
			c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityExistMessage(customerTitle)})
		} else {
			uid := models.CreateCustomer(w)
			c.JSON(http.StatusOK, ApiStatus{uid,utils.SuccessMessage})
		}
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func ReadCustomerById(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)
	w := models.GetCustomerById(id)
	if w != nil {
		c.JSON(http.StatusOK, w)
	} else {
		c.JSON(http.StatusNotFound, ApiMessage{utils.EntityNotExistMessage(customerTitle)})
	}
}

func UpdateCustomer(c *gin.Context) {
	var w models.Customer
	if err := c.ShouldBindJSON(&w); err == nil {
		existItem := models.GetCustomerById(w.Id)
		if existItem == nil {
			c.JSON(http.StatusUnprocessableEntity, ApiStatus{existItem.Id,utils.EntityWithIdNotExistMessage(customerTitle)})
		}else {
			if models.UpdateCustomer(w) {
				c.JSON(http.StatusOK, ApiStatus{w.Id,utils.SuccessMessage})
			} else {
				c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotUpdatedMessage(customerTitle)})
			}
		}
	} else {
		c.AbortWithError(http.StatusBadRequest, err)
	}
}

func DeleteCustomer(c *gin.Context) {
	idParam := c.Param(utils.IdKey)
	id, err := strconv.ParseInt(idParam, 10, 64)
	utils.CheckErr(err)

	existItem := models.GetCustomerById(id)
	if existItem == nil {
		c.JSON(http.StatusUnprocessableEntity, ApiMessage{utils.EntityNotExistMessage(customerTitle)})
	} else {
		if models.DeleteCustomer(*existItem) {
			c.JSON(http.StatusOK, ApiMessage{utils.SuccessMessage})
		} else {
			c.JSON(http.StatusBadRequest, ApiMessage{utils.EntityNotDeletedMessage(customerTitle)})
		}
	}
}

func GetCustomers(c *gin.Context) {
	townParam := c.Param("town")
	town, err := strconv.ParseInt(townParam, 10, 64)
	utils.CheckErr(err)

	c.JSON(http.StatusOK, models.GetCustomerByTown(town))

}
