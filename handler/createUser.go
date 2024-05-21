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

	// Create a Firestore context
	ctx := context.Background()

	// Bind the request body to the user struct
	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Check required fields
	if user.Uid == "" {
		errParamIsRequired(c, "uid", "string")
		return
	}

	if user.UserDisplayName == "" {
		errParamIsRequired(c, "userDisplayName", "string")
		return
	}

	if user.UserEmail == "" {
		errParamIsRequired(c, "userEmail", "string")
		return
	}

	// Set timestamps and default values
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActivated = true
	user.Id = user.Uid

	// Check if the user already exists
	_, err := dbClient.Collection("users").Select("uid").Where("uid", "==", user.Uid).Documents(ctx).Next()

	if err == nil {
		sendError(c, http.StatusConflict, "user already exists")
		return
	}

	// Add the user to Firestore
	_, _, err = dbClient.Collection("users").Add(ctx, user)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user created", http.StatusCreated, user)
}
