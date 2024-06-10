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

	var userHandler *handler.AuthHandler

	var adminHandler *handler.AuthHandler

	var productHandler *handler.AuthHandler

	var cartHandler *handler.AuthHandler

	var checkoutHandler *handler.AuthHandler

	userCollection := db.Collection("users")

	adminCollection := db.Collection("admin")

	productCollection := db.Collection("products")

	cartCollection := db.Collection("carts")

	checkoutCollection := db.Collection("checkouts")

	authHandler = handler.NewAuthHandler(userCollection, authClient, ctx)

	userHandler = handler.NewAuthHandler(userCollection, authClient, ctx)

	adminHandler = handler.NewAuthHandler(adminCollection, authClient, ctx)

	productHandler = handler.NewAuthHandler(productCollection, authClient, ctx)

	cartHandler = handler.NewAuthHandler(cartCollection, authClient, ctx)

	checkoutHandler = handler.NewAuthHandler(checkoutCollection, authClient, ctx)

	UserHandler := handler.NewDefaultHandler(userCollection, ctx)

	AdminHandler := handler.NewDefaultHandler(adminCollection, ctx)

	ProductHandler := handler.NewDefaultHandler(productCollection, ctx)

	CartHandler := handler.NewDefaultHandler(cartCollection, ctx)

	CheckoutHandler := handler.NewDefaultHandler(checkoutCollection, ctx)

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

	userAuthorized := router.Group("api/v1/users/")

	adminAuthorized := router.Group("api/v1/admin/")

	productAuthorized := router.Group("api/v1/products/")

	cartAuthorized := router.Group("api/v1/carts/")

	checkoutAuthorized := router.Group("api/v1/checkout/")

	authorized.Use(authHandler.AuthMiddleware())

	userAuthorized.Use(userHandler.AuthMiddleware())

	adminAuthorized.Use(adminHandler.AuthMiddleware())

	productAuthorized.Use(productHandler.AuthMiddleware())

	cartAuthorized.Use(cartHandler.AuthMiddleware())

	checkoutAuthorized.Use(checkoutHandler.AuthMiddleware())

	router.GET("api/v1/predictions", handler.GetPredictions)
	router.GET("api/v1/products/:id", handler.GetProduct)
	router.GET("api/v1/products", handler.GetAllProduct)

	{
		authorized.GET("/ping", UserHandler.Ping)
		userAuthorized.POST("/", UserHandler.CreateUser)
		userAuthorized.GET("/:id", UserHandler.GetUser)
		userAuthorized.PUT("/:id", UserHandler.UpdateUser)
		userAuthorized.DELETE("/:id", UserHandler.DeleteUser)
		adminAuthorized.GET("/:id", AdminHandler.GetAdmin)
		adminAuthorized.POST("/", AdminHandler.CreateAdmin)
		adminAuthorized.PUT("/:id", AdminHandler.UpdateAdmin)
		adminAuthorized.DELETE("/:id", AdminHandler.DeleteAdmin)

		productAuthorized.POST("/", ProductHandler.CreateProduct)
		productAuthorized.PUT("/:id", ProductHandler.UpdateProduct)
		productAuthorized.DELETE("/:id", ProductHandler.DeleteProduct)
		productAuthorized.PUT("/all", ProductHandler.UpdateMultipleProducts)
		productAuthorized.POST("/all", ProductHandler.UpdateOrAddMultipleProducts)

		cartAuthorized.GET("/", CartHandler.GetCart)
		cartAuthorized.PUT("/", CartHandler.CreateOrUpdateCart)

		checkoutAuthorized.GET("/", CheckoutHandler.GetCheckout)
		checkoutAuthorized.POST("/", CheckoutHandler.CreateCheckout)

	}

}
