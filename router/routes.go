package router

import (
	"context"
	"os"

	"github.com/GinoCodeSpace/bridge/handler"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	ctx := context.Background()

	db, authClient := handler.InitializeHandler()

	var authHandler *handler.AuthHandler

	var adminHandler *handler.AuthHandler

	var productHandler *handler.AuthHandler

	userCollection := db.Collection("users")

	adminCollection := db.Collection("admin")

	productCollection := db.Collection("products")

	authHandler = handler.NewAuthHandler(userCollection, authClient, ctx)

	adminHandler = handler.NewAuthHandler(adminCollection, authClient, ctx)

	productHandler = handler.NewAuthHandler(productCollection, authClient, ctx)

	UserHandler := handler.NewDefaultHandler(userCollection, ctx)

	AdminHandler := handler.NewDefaultHandler(adminCollection, ctx)

	ProductHandler := handler.NewDefaultHandler(productCollection, ctx)

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

	authorized := router.Group("api/v1/")

	adminAuthorized := router.Group("api/v1/admin/")

	productAuthorized := router.Group("api/v1/products/")

	authorized.Use(authHandler.AuthMiddleware())

	adminAuthorized.Use(adminHandler.AuthMiddleware())

	productAuthorized.Use(productHandler.AuthMiddleware())

	router.GET("api/v1/predictions", handler.GetPredictions)
	router.GET("api/v1/products/:id", handler.GetProduct)

	{
		authorized.GET("/ping", UserHandler.Ping)
		authorized.POST("/users", UserHandler.CreateUser)
		authorized.GET("/users/:id", UserHandler.GetUser)
		authorized.PUT("/users/:id", UserHandler.UpdateUser)
		authorized.DELETE("/users/:id", UserHandler.DeleteUser)
		adminAuthorized.GET("/:id", AdminHandler.GetAdmin)
		adminAuthorized.POST("/", AdminHandler.CreateAdmin)
		adminAuthorized.PUT("/:id", AdminHandler.UpdateAdmin)
		adminAuthorized.DELETE("/:id", AdminHandler.DeleteAdmin)

		productAuthorized.POST("/", ProductHandler.CreateProduct)
		productAuthorized.PUT("/:id", ProductHandler.UpdateProduct)
		productAuthorized.DELETE("/:id", ProductHandler.DeleteProduct)
	}

}
