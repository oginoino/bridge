package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) CreateProduct(c *gin.Context) {
	var product models.Product
	ctx := context.Background()

	uid, _ := c.Get("uid")
	var admin models.AdminUser

	documentUserSnapShot, err := dbClient.Collection("admin").Where("uid", "==", uid).Documents(ctx).Next()

	if err != nil {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return
	}

	documentUserSnapShot.DataTo(&admin)

	if !admin.IsActivated {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return
	}

	if admin.Role != "admin" {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return
	}

	if err := c.ShouldBindJSON(&product); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	if product.ProductCode == "" {
		sendError(c, http.StatusBadRequest, "product code is required")
		return
	}

	if product.ProductName == "" {
		sendError(c, http.StatusBadRequest, "product name is required")
		return
	}

	if product.ProductPrice == 0 {
		sendError(c, http.StatusBadRequest, "product price is required")
		return
	}

	if product.ProductUnitOfMeasure == "" {
		sendError(c, http.StatusBadRequest, "product unit of measure is required")
		return
	}

	if product.ProductUnitQuantity == "" {
		sendError(c, http.StatusBadRequest, "product unit quantity is required")
		return
	}

	if len(product.ProductCategories) == 0 {
		sendError(c, http.StatusBadRequest, "product categories is required")
		return
	}

	if product.AvailableQuantity == 0 {
		sendError(c, http.StatusBadRequest, "available quantity is required")
		return
	}

	product.Id = product.ProductCode

	documentSnapShot, err := dbClient.Collection(handler.collection.ID).Where("productCode", "==", product.ProductCode).Documents(ctx).Next()

	if documentSnapShot != nil || err == nil {
		sendError(c, http.StatusConflict, "product already exists")
		return
	}

	_, err = dbClient.Collection(handler.collection.ID).Doc(product.Id).Set(ctx, product)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(c, "product created successfully", http.StatusCreated, product)

}
