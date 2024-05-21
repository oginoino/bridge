package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetPredictions(c *gin.Context) {
	input := c.Query("input")
	if input == "" {
		sendError(c, http.StatusBadRequest, "input is required")
		return
	}

	mapsService := NewMapsService()
	predictions, err := mapsService.FetchAddress(input)
	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, predictions)
}
