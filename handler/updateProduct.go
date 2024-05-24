package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	var product models.Product
	var existingProduct models.Product

	isAdmin := CheckAdminFunction(ctx, c)

	if isAdmin {
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	query := dbClient.Collection(handler.collection.ID).Where("id", "==", id).Limit(1)
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

	doc.DataTo(&existingProduct)

	product.Id = existingProduct.Id
	product.ProductCode = existingProduct.ProductCode

	if product.ProductName == "" {
		product.ProductName = existingProduct.ProductName
	}

	if product.ProductPrice == 0 {
		product.ProductPrice = existingProduct.ProductPrice
	}

	if product.ProductUnitOfMeasure == "" {
		product.ProductUnitOfMeasure = existingProduct.ProductUnitOfMeasure
	}

	if product.ProductUnitQuantity == "" {
		product.ProductUnitQuantity = existingProduct.ProductUnitQuantity
	}

	if len(product.ProductCategories) == 0 {
		product.ProductCategories = existingProduct.ProductCategories
	}

	if product.AvailableQuantity == 0 {
		product.AvailableQuantity = existingProduct.AvailableQuantity
	}

	if product.ContentValue == "" {
		product.ContentValue = existingProduct.ContentValue
	}

	if product.ProductImageSrc == "" {
		product.ProductImageSrc = existingProduct.ProductImageSrc
	}

	if product.ProducKilogramsWeight == 0 {
		product.ProducKilogramsWeight = existingProduct.ProducKilogramsWeight
	}

	if product.ProductCubicMeterVolume == 0 {
		product.ProductCubicMeterVolume = existingProduct.ProductCubicMeterVolume
	}

	if !product.IsActivated {
		product.IsActivated = existingProduct.IsActivated
	}

	_, err = dbClient.Collection(handler.collection.ID).Doc(id).Set(ctx, product)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, product)

}
