package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) CreateCheckout(c *gin.Context) {
	ctx := context.Background()

	var checkout models.Checkout
	var user models.User
	var existingCart models.Cart

	if err := c.ShouldBindJSON(&checkout); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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

	checkout.Payer = user

	queryCart := dbClient.Collection(handler.collection.ID).Where("shopper.uid", "==", uid).Limit(1)

	docsCart, err := queryCart.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docsCart) > 0 {
		doc := docsCart[0]
		doc.DataTo(&existingCart)
		checkout.ProductCart = existingCart
	} else {
		sendError(c, http.StatusNotFound, "Cart not found")
		return
	}

	checkout.CheckoutId = handler.collection.NewDoc().ID
	checkout.CreartedAt = time.Now()
	checkout.UpdatedAt = time.Now()
	checkout.IsTimerStarted = true
	checkout.DeliveryFee = 5.0
	checkout.Payment.Status = "pending"
	checkout.DeliveryTime = 15
	checkout.RemainingSeconds = 300
	checkout.IsTimerStarted = true

	_, err = dbClient.Collection(handler.collection.ID).Doc(checkout.CheckoutId).Set(ctx, checkout)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "Checkout criado com sucesso!", http.StatusCreated, checkout)

}
