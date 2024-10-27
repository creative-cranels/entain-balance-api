package server

import (
	"balance-api/handler"
	"balance-api/repository"
	"balance-api/service"
	"balance-api/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigureRoutes(server *Server) {
	server.Gin.Use(corsMiddleware)
	server.Gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Repository Initialization
	userRepo := repository.NewUserRepository(server.DB)
	transactionRepo := repository.NewTransctionRepository(server.DB)

	// Services Initialization
	userService := service.NewUserService(userRepo, transactionRepo)

	// Handlers Initialization
	welcomeHandler := handler.NewWelcomeHandler()
	userHandler := handler.NewUserHandler(userService)

	// Public routes
	publicRoutes := server.Gin.Group("")
	publicRoutes.GET("/ping", welcomeHandler.Ping)

	userRoutes := server.Gin.Group("/user", CreateRequestWrapper)
	userRoutes.GET("/:id/balance", userHandler.GetUserBalance)
	userRoutes.POST("/:id/transaction", userHandler.MakeTransaction)
}

func corsMiddleware(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", ctx.GetHeader("Origin"))
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set(
		"Access-Control-Allow-Headers",
		`Accept-Language,
		Content-Type,
		Content-Length,
		Accept-Encoding,
		X-CSRF-Token,
		Authorization,
		accept,
		origin,
		Cache-Control,
		X-Requested-With,
		grant_type,
		Grant-Type,
		Accept,
		Referer,
		User-Agent,
		Trace-Id`)
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

	if ctx.Request.Method == "OPTIONS" {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}
	ctx.Next()
}

func CreateRequestWrapper(ctx *gin.Context) {
	wrapper := utils.RequestWrapper{
		C:              ctx,
		ID:             0,
		Q:              "",
		Page:           -1,
		PerPage:        -1,
		DefaultPage:    0,
		DefaultPerPage: 20,
	}
	wrapper.ParseDefaultQueryParams()

	ctx.Set("request-wrapper", wrapper)
	ctx.Next()
}
