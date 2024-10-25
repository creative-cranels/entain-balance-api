package server

import (
	"balance-api/config"
	"balance-api/db"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Cfg *config.Config
	Gin *gin.Engine
	DB  *db.Database
}

func NewServer(cfg *config.Config) *Server {
	return &Server{
		Cfg: cfg,
		Gin: gin.Default(),
		DB:  db.InitDB(cfg.DB),
	}
}

func (server *Server) Run(addr string) error {
	return server.Gin.Run(":" + addr)
}
