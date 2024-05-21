package router

import (
	"context"

	"github.com/GinoCodeSpace/bridge/handler"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	ctx := context.Background()

	db, authClient := handler.InitializeHandler()

	var authHandler *handler.AuthHandler

	userCollection := db.Collection("users")

	authHandler = handler.NewAuthHandler(userCollection, authClient, ctx)

	UserHandler := handler.NewDefaultHandler(userCollection, ctx)

	corsConfig := cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	authorized := router.Group("api/v1/")

	authorized.Use(authHandler.AuthMiddleware())

	{
		authorized.GET("/ping", UserHandler.Ping)
	}
}
