package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) UpdateOrAddMultipleProducts(c *gin.Context) {
	ctx := context.Background()

	var products []models.Product

	isAdmin := CheckAdminFunction(ctx, c)

	if isAdmin {
		return
	}

	if err := c.ShouldBindJSON(&products); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	updatedOrAddedProducts := make([]models.Product, 0)

	for _, product := range products {
		var existingProduct models.Product
		query := dbClient.Collection(handler.collection.ID).Where("productCode", "==", product.ProductCode).Limit(1)
		docs, err := query.Documents(ctx).GetAll()

		if err != nil {
			sendError(c, http.StatusInternalServerError, err.Error())
			return
		}

		if len(docs) > 0 {
			doc := docs[0]
			doc.DataTo(&existingProduct)

			product.Id = existingProduct.Id

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

			_, err = dbClient.Collection(handler.collection.ID).Doc(existingProduct.Id).Set(ctx, product)

			if err != nil {
				sendError(c, http.StatusInternalServerError, err.Error())
				return
			}

		} else {
			// Adiciona novo produto
			docRef, _, err := dbClient.Collection(handler.collection.ID).Add(ctx, product)

			if err != nil {
				sendError(c, http.StatusInternalServerError, err.Error())
				return
			}

			product.Id = docRef.ID
		}

		updatedOrAddedProducts = append(updatedOrAddedProducts, product)
	}

	c.JSON(http.StatusOK, updatedOrAddedProducts)
}
