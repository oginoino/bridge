package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) GetCart(c *gin.Context) {
	ctx := context.Background()

	uid, _ := c.Get("uid")

	var cart models.Cart

	queryCart := dbClient.Collection(handler.collection.ID).Where("shopper.uid", "==", uid).Limit(1)

	docsCart, err := queryCart.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docsCart) == 0 {
		sendError(c, http.StatusNotFound, "Cart not found")
		return
	}

	doc := docsCart[0]

	doc.DataTo(&cart)

	sendSuccess(c, "Get cart with success!", http.StatusOK, cart)

}
