package router

import (
	"gin-basic/internal/database"
	"gin-basic/internal/handler"
	"gin-basic/internal/repository"
	"gin-basic/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthRoute(rg *gin.RouterGroup, db *database.DB) {

	// Initialize dependencies
	autoRepository := repository.NewRepository(db)
	authService := service.NewAuthService(*autoRepository)
	authHandler := handler.NewAuthHandler(*authService)

	authRouter := rg.Group("/auth")
	{
		authRouter.GET("/login", authHandler.LoaginHandler)
		authRouter.GET("/login/:userId", authHandler.LoaginHandler)
		authRouter.GET("/reg")
	}
}
