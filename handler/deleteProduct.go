package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	ctx := context.Background()

	isAdmin := CheckAdminFunction(ctx, c)

	if isAdmin {
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

	_, err = doc.Ref.Delete(ctx)

	if err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Product deleted successfully",
	})
}
