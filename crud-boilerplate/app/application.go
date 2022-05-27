package app

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/andrewjoyt/crud-boilerplate/database/connection"
)

var (
	router = gin.Default()
)

func StartApplication() {

	mapUrls()

	connection.Connect()

	port := os.Getenv("APP_PORT")
	fmt.Printf("running at port: %s", port)
	router.Run(port)
}
