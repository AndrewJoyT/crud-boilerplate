package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	regisDomain "github.com/AndrewJoyT/crud-boilerplate/domain/register"
	"github.com/AndrewJoyT/crud-boilerplate/services"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"
)

func RegisterUser(c *gin.Context) {
	var requestBody regisDomain.RegisterUserRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errInsert := services.RegisterService.CreateUser(&requestBody)
	if errInsert != nil {
		c.JSON(errInsert.Status, &errors.RestErr{
			Message: errInsert.Message,
			Status:  errInsert.Status,
			Error:   errInsert.Error,
		})
		return
	}
	c.JSON(http.StatusCreated, result)
}

func RegisterConfirmation(c *gin.Context) {
	token := c.Param("token")
	errInsert := services.RegisterService.Confirmation(token)
	if errInsert != nil {
		c.JSON(errInsert.Status, &errors.RestErr{
			Message: errInsert.Message,
			Status:  errInsert.Status,
			Error:   errInsert.Error,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
