package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AndrewJoyT/crud-boilerplate/services"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"

	loginDomain "github.com/AndrewJoyT/crud-boilerplate/domain/login"
	"github.com/AndrewJoyT/crud-boilerplate/domain/response"
)

func Login(c *gin.Context) {
	var requestBody loginDomain.LoginRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errLogin := services.LoginService.Login(&requestBody)
	if errLogin != nil {
		c.JSON(errLogin.Status, &errors.RestErr{
			Message: errLogin.Message,
			Status:  errLogin.Status,
			Error:   errLogin.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully login"))
}
