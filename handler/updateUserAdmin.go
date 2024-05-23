package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) UpdateAdmin(c *gin.Context) {
	id := c.Param("id")

	var admin models.AdminUser
	var existingAdminUser models.AdminUser

	ctx := context.Background()

	if err := c.ShouldBindJSON(&admin); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if admin.Uid == "" || admin.UserDisplayName == "" || admin.UserEmail == "" {
		sendError(c, http.StatusBadRequest, "uid, userDisplayName, and userEmail are required fields")
		return
	}

	query := dbClient.Collection(handler.collection.ID).Where("id", "==", id).Limit(1)
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
	doc.DataTo(&existingAdminUser)

	if admin.UserDisplayName != "" {
		existingAdminUser.UserDisplayName = admin.UserDisplayName
	}
	if admin.UserEmail != "" {
		existingAdminUser.UserEmail = admin.UserEmail
	}
	if admin.UserPhotoUrl != "" {
		existingAdminUser.UserPhotoUrl = admin.UserPhotoUrl
	}
	if admin.UserName != "" {
		existingAdminUser.UserName = admin.UserName
	}
	if len(admin.UserProperties) > 0 {
		existingAdminUser.UserProperties = admin.UserProperties
	}
	if len(admin.Addresses) > 0 {
		existingAdminUser.Addresses = admin.Addresses
	}
	if (admin.SelectedAddress != models.Address{}) {
		existingAdminUser.SelectedAddress = admin.SelectedAddress
	}

	if admin.Role != "" {
		existingAdminUser.Role = admin.Role
	}

	existingAdminUser.UpdatedAt = models.CustomTime{Time: time.Now()}

	_, err = dbClient.Collection(handler.collection.ID).Doc(doc.Ref.ID).Set(ctx, existingAdminUser)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user updated", http.StatusOK, existingAdminUser)
}
