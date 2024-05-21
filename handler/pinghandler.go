package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
