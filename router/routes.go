package router

import (
	"github.com/GinoCodeSpace/bridge/handler"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// ctx := context.Background()

	// db, authClient := handler.InitializeHandler()

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	authorized := router.Group("api/v1/")

	var authHandler *handler.AuthHandler

	authorized.Use(authHandler.AuthMiddleware())

	// router.GET("/ping", handler.Ping)

	{
		authorized.GET("/ping", handler.Ping)
	}
}
