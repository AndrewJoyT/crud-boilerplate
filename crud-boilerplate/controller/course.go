package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/AndrewJoyT/crud-boilerplate/services"
	"github.com/AndrewJoyT/crud-boilerplate/utils/errors"

	courseDomain "github.com/AndrewJoyT/crud-boilerplate/domain/course"
	"github.com/AndrewJoyT/crud-boilerplate/domain/response"
)

func CourseFileUpload(c *gin.Context) {
	var requestBody courseDomain.CourseFileUploadRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errUpload := services.CourseService.CourseFileUpload(&requestBody)
	if errUpload != nil {
		c.JSON(errUpload.Status, &errors.RestErr{
			Message: errUpload.Message,
			Status:  errUpload.Status,
			Error:   errUpload.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully added file"))
}

func CreateCourse(c *gin.Context) {
	var requestBody courseDomain.CreateCourseRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}
	result, errUpload := services.CourseService.CreateCourse(&requestBody)
	if errUpload != nil {
		c.JSON(errUpload.Status, &errors.RestErr{
			Message: errUpload.Message,
			Status:  errUpload.Status,
			Error:   errUpload.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully added file"))
}

func GetCourseByID(c *gin.Context) {
	id := c.Param("id")
	result, errUpload := services.CourseService.GetCourseByID(id)
	if errUpload != nil {
		c.JSON(errUpload.Status, &errors.RestErr{
			Message: errUpload.Message,
			Status:  errUpload.Status,
			Error:   errUpload.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully added file"))
}

func JoinCourse(c *gin.Context) {
	var requestBody courseDomain.JoinCourseRequest
	err := c.ShouldBindJSON(&requestBody)
	if err != nil {
		c.JSON(http.StatusBadRequest, &errors.RestErr{
			Message: "invalid json body",
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
		})
		return
	}

	result, errUpload := services.CourseService.JoinCourse(&requestBody)
	if errUpload != nil {
		c.JSON(errUpload.Status, &errors.RestErr{
			Message: errUpload.Message,
			Status:  errUpload.Status,
			Error:   errUpload.Error,
		})
		return
	}
	c.JSON(http.StatusOK, response.SuccessResponseMap(result, "succesfully added file"))
}
