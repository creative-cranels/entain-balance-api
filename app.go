package application

import (
	"balance-api/config"
	"balance-api/server"
	"log"

	"github.com/gin-gonic/gin"
)

func Start(cfg *config.Config) {
	gin.SetMode(cfg.HTTP.Mode)

	app := server.NewServer(cfg)

	server.ConfigureRoutes(app)

	err := app.Run(cfg.HTTP.Port)
	if err != nil {
		log.Fatal("Port already used")
	}
}
