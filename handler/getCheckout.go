package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) GetCheckout(c *gin.Context) {
	ctx := context.Background()

	uid, _ := c.Get("uid")

	var checkout models.Checkout

	queryCheckout := dbClient.Collection(handler.collection.ID).Where("payer.uid", "==", uid).Where("payment.status", "==", "pending").Limit(1)

	docsCkeckout, err := queryCheckout.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docsCkeckout) == 0 {
		sendError(c, http.StatusNotFound, "Checkout not found")
		return
	}

	doc := docsCkeckout[0]

	doc.DataTo(&checkout)

	sendSuccess(c, "Get checkout with success!", http.StatusOK, checkout)

}
