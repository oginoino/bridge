package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

func (handler *DefaultHandler) CreateUser(c *gin.Context) {

	var user models.User

	ctx := context.Background()

	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the user struct
	if err := handler.validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": errorMessages})
		return
	}

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.IsActivated = true
	user.Id = user.Uid

	_, err := dbClient.Collection("users").Select("uid").Where("uid", "==", user.Uid).Documents(ctx).Next()

	if err == nil {
		sendError(c, http.StatusConflict, "user already exists")
		return
	}

	_, _, err = dbClient.Collection("users").Add(ctx, user)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user created", http.StatusCreated, user)
}
