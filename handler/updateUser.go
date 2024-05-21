package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func (handler *DefaultHandler) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	collectionName := "users"

	var user models.User

	ctx := context.Background()

	if err := c.ShouldBindJSON(&user); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	// Validate the user struct
	if err := validate.Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		errorMessages := make(map[string]string)
		for _, fieldError := range validationErrors {
			errorMessages[fieldError.Field()] = fieldError.Tag()
		}
		c.JSON(http.StatusBadRequest, gin.H{"validationErrors": errorMessages})
		return
	}

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

	// Use the first document found (there should be only one)
	doc := docs[0]

	user.UpdatedAt = time.Now()
	user.Id = id
	user.Uid = id
	user.CreatedAt = doc.Data()["createdAt"].(time.Time)

	_, err = dbClient.Collection(collectionName).Doc(doc.Ref.ID).Set(ctx, user)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "user updated", http.StatusOK, user)
}
