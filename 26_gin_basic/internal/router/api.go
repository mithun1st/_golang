package router

import (
	"gin-basic/internal/database"
	"gin-basic/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoute(db *database.DB) *gin.Engine {

	var router *gin.Engine = gin.Default()

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			AuthRoute(v1, db)
		}

		v2 := api.Group("/v2")
		v2.Use(middleware.AuthToken)
		{
			AuthRoute(v2, db)
		}

	}

	return router
}
