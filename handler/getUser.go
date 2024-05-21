package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) GetUser(c *gin.Context) {
	id := c.Param("id")
	collectionName := "users"
	var user = models.User{}

	ctx := context.Background()

	query := dbClient.Collection(collectionName).Where("id", "==", id).Limit(1)
	snapshot, err := query.Documents(ctx).Next()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if !snapshot.Exists() {
		sendError(c, http.StatusNotFound, "User not found")
		return
	}

	snapshot.DataTo(&user)

	c.JSON(http.StatusOK, user)

}
