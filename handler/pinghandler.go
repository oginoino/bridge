package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {

	id := c.Query("id")

	if id == "" {
		sendError(c, http.StatusBadRequest, errParamIsRequired("id", "string").Error())
		return
	}

	sendSuccess(c, "Poing", http.StatusOK, gin.H{"id": id})
}
