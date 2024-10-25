package main

import (
	application "balance-api"
	"balance-api/config"
	"balance-api/docs"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// @title Balance API
// @version 1.0
// @description Balance API for manipulating user's balance using transactions

// @contact.name Aza M
// @contact.email mukhamejanov.aza@gmail.com

// @BasePath /
func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", os.Getenv("HOST"), os.Getenv("PORT"))
	application.Start(config.NewConfig())
}
