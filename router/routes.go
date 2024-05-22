package router

import (
	"context"
	"os"

	"github.com/GinoCodeSpace/bridge/handler"
	"github.com/gin-contrib/cors"
	"gopkg.in/go-playground/validator.v9"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	ctx := context.Background()

	db, authClient := handler.InitializeHandler()

	var authHandler *handler.AuthHandler

	var validate = validator.New()

	userCollection := db.Collection("users")

	authHandler = handler.NewAuthHandler(userCollection, authClient, ctx)

	UserHandler := handler.NewDefaultHandler(userCollection, ctx, validate)

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "*"
	}

	corsConfig := cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Access-Control-Allow-Headers", "Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}

	router.Use(cors.New(corsConfig))

	router.GET("/predictions", handler.GetPredictions)

	authorized := router.Group("api/v1/")

	authorized.Use(authHandler.AuthMiddleware())

	{
		authorized.GET("/ping", UserHandler.Ping)
		authorized.POST("/users", UserHandler.CreateUser)
		authorized.GET("/users/:id", UserHandler.GetUser)
		authorized.PUT("/users/:id", UserHandler.UpdateUser)
		authorized.DELETE("/users/:id", UserHandler.DeleteUser)
	}

}
