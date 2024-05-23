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

	uid, _ := c.Get("uid")
	admin.Uid = uid.(string)
	admin.UserDisplayName = c.GetString("name")
	admin.UserEmail = c.GetString("email")
	admin.Id = admin.Uid
	admin.CreatedAt = models.CustomTime{Time: time.Now()}
	admin.UpdatedAt = models.CustomTime{Time: time.Now()}
	admin.IsActivated = true
	admin.Role = "admin"

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
