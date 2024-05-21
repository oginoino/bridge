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

	// Bind the request body to the user struct

	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Check required fields
	if user.Uid == "" {
		errParamIsRequired("uid", "uid is required")
		return
	}

	if user.UserDisplayName == "" {
		errParamIsRequired("userDisplayName", "userDisplayName is required")
		return
	}

	if user.UserEmail == "" {
		errParamIsRequired("userEmail", "userEmail is required")
		return
	}

	// Set timestamps and default values
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActivated = true

	// Create a Firestore context
	ctx := context.Background()

	// Add the user to Firestore
	_, _, err := dbClient.Collection("users").Add(ctx, user)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user created", http.StatusCreated, user)
}
