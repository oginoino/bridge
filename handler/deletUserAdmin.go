package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *DefaultHandler) DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	query := dbClient.Collection(handler.collection.ID).Where("id", "==", id).Limit(1)
	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		sendError(c, 500, err.Error()+" "+id)
		return
	}

	if len(docs) == 0 {
		sendError(c, 404, "User not found")
		return
	}

	doc := docs[0]

	_, err = doc.Ref.Delete(ctx)
	if err != nil {
		sendError(c, 500, err.Error()+" "+id)
		return
	}

	sendSuccess(c, "User deleted", http.StatusOK, gin.H{
		"id":      id,
		"status":  "success",
		"message": "User deleted",
		"data":    nil,
	})

}
