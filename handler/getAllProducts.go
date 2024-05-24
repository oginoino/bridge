package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func GetAllProduct(c *gin.Context) {
	ctx := context.Background()

	collectionRef := "products"

	query := dbClient.Collection(collectionRef)

	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	var products []models.Product

	for _, doc := range docs {
		var product models.Product

		if err := doc.DataTo(&product); err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		products = append(products, product)
	}

	c.JSON(http.StatusOK, products)
}
