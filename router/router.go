package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// Initialize routes
	initializeRoutes(router)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)

}
