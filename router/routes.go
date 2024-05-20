package router

import (
	"github.com/GinoCodeSpace/bridge/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	router.GET("/ping", handler.Ping)
}
