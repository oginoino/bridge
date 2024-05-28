package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) CreateOrUpdateCart(c *gin.Context) {
	ctx := context.Background()

	var cart models.Cart
	var existingCart models.Cart
	var user models.User

	if err := c.ShouldBindJSON(&cart); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	uid, _ := c.Get("uid")

	queryUser := dbClient.Collection("users").Where("uid", "==", uid).Limit(1)

	docsUser, err := queryUser.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docsUser) == 0 {
		sendError(c, http.StatusNotFound, "User not found")
		return
	}

	doc := docsUser[0]
	doc.DataTo(&user)

	cart.Shopper = user

	queryCart := dbClient.Collection(handler.collection.ID).Where("shopper.uid", "==", uid).Limit(1)

	docsCart, err := queryCart.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docsCart) > 0 {
		doc := docsCart[0]
		doc.DataTo(&existingCart)
		cart.ID = existingCart.ID
	} else {
		newDocRef := dbClient.Collection(handler.collection.ID).NewDoc()
		cart.ID = newDocRef.ID
	}

	_, err = dbClient.Collection(handler.collection.ID).Doc(cart.ID).Set(ctx, cart)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "Cart atualizado com sucesso!", http.StatusOK, cart)
}
