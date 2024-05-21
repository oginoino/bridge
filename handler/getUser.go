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
	var user models.User

	ctx := context.Background()

	query := dbClient.Collection(collectionName).Where("id", "==", id).Limit(1)
	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docs) == 0 {
		sendError(c, http.StatusNotFound, "User not found")
		return
	}

	doc := docs[0]
	if err := doc.DataTo(&user); err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
