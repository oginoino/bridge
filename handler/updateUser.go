package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user models.User
	var existingUser models.User

	ctx := context.Background()

	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if user.Uid == "" || user.UserDisplayName == "" || user.UserEmail == "" {
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
	doc.DataTo(&existingUser)

	if user.UserDisplayName != "" {
		existingUser.UserDisplayName = user.UserDisplayName
	}
	if user.UserEmail != "" {
		existingUser.UserEmail = user.UserEmail
	}
	if user.UserPhotoUrl != "" {
		existingUser.UserPhotoUrl = user.UserPhotoUrl
	}
	if user.UserName != "" {
		existingUser.UserName = user.UserName
	}
	if len(user.UserProperties) > 0 {
		existingUser.UserProperties = user.UserProperties
	}
	if len(user.Addresses) > 0 {
		existingUser.Addresses = user.Addresses
	}
	if (user.SelectedAddress != models.Address{}) {
		existingUser.SelectedAddress = user.SelectedAddress
	}
	existingUser.UpdatedAt = models.CustomTime{Time: time.Now()}

	_, err = dbClient.Collection(handler.collection.ID).Doc(doc.Ref.ID).Set(ctx, existingUser)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user updated", http.StatusOK, existingUser)
}
