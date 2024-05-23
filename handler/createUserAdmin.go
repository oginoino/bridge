package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) CreateAdmin(c *gin.Context) {
	var admin models.AdminUser
	ctx := context.Background()

	if err := c.ShouldBindJSON(&admin); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if admin.Uid == "" || admin.UserDisplayName == "" || admin.UserEmail == "" || admin.Role == "" {
		sendError(c, http.StatusBadRequest, "uid, userDisplayName, userEmail, and role are required fields")
		return
	}

	admin.Id = admin.Uid
	admin.CreatedAt = models.CustomTime{Time: time.Now()}
	admin.UpdatedAt = models.CustomTime{Time: time.Now()}
	admin.IsActivated = true

	_, err := dbClient.Collection(handler.collection.ID).Where("uid", "==", admin.Uid).Documents(ctx).Next()

	if err == nil {
		sendError(c, http.StatusConflict, "user already exists")
		return
	}

	_, _, err = dbClient.Collection(handler.collection.ID).Add(ctx, admin)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "admin created", http.StatusCreated, admin)
}
