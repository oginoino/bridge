package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	ctx := context.Background()

	collectionRef := "products"

	query := dbClient.Collection(collectionRef).Where("id", "==", id).Limit(1)

	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if len(docs) == 0 {
		sendError(c, http.StatusNotFound, "Product not found")
		return
	}

	doc := docs[0]

	var product models.Product

	if err := doc.DataTo(&product); err != nil {

		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)

}
