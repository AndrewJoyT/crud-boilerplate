package app

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/AndrewJoyT/crud-boilerplate/controller"
)

func mapUrls() {

	v1 := router.Group("api/v1")
	{
		userPath := v1.Group("/user", gin.BasicAuth(gin.Accounts{
			os.Getenv("SERVICE_USERNAME"): os.Getenv("SERVICE_PASSWORD"),
		}))
		{
			registerPath := userPath.Group("/register")
			{
				registerPath.POST("/", controller.RegisterUser)
				registerPath.GET("/confirmation/:token", controller.RegisterConfirmation)
				//registerPath.POST("/teacher", controller.RegisterAdmin)
				//registerPath.POST("/student", controller.RegisterAdmin)
			}
			loginPath := userPath.Group("/login")
			{
				loginPath.POST("/", controller.Login)
			}
		}
		coursePath := v1.Group("/course", gin.BasicAuth(gin.Accounts{
			os.Getenv("SERVICE_USERNAME"): os.Getenv("SERVICE_PASSWORD"),
		}))
		{
			coursePath.POST("/", controller.CreateCourse)
			coursePath.GET("/:id", controller.GetCourseByID)
			coursePath.POST("/upload", controller.CourseFileUpload)
			coursePath.POST("/join", controller.JoinCourse)
		}
	}

	//router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
