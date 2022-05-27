package main

import (
	"github.com/joho/godotenv"
	"github.com/AndrewJoyT/crud-boilerplate/app"
)

func main() {
	_ = godotenv.Load()

	app.StartApplication()
}
