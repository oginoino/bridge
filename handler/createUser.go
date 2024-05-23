package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) CreateUser(c *gin.Context) {
	var user models.User
	ctx := context.Background()

	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.Uid == "" || user.UserDisplayName == "" || user.UserEmail == "" {
		sendError(c, http.StatusBadRequest, "uid, userDisplayName, and userEmail are required fields")
		return
	}

	user.Id = user.Uid
	user.CreatedAt = models.CustomTime{Time: time.Now()}
	user.UpdatedAt = models.CustomTime{Time: time.Now()}
	user.IsActivated = true

	_, err := dbClient.Collection(handler.collection.ID).Where("uid", "==", user.Uid).Documents(ctx).Next()

	if err == nil {
		sendError(c, http.StatusConflict, "user already exists")
		return
	}

	_, _, err = dbClient.Collection(handler.collection.ID).Add(ctx, user)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user created", http.StatusCreated, user)
}
