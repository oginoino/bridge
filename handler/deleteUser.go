package handler

import "github.com/gin-gonic/gin"

func (handler *DefaultHandler) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	query := dbClient.Collection(handler.collection.ID).Where("id", "==", id).Limit(1)
	docs, err := query.Documents(ctx).GetAll()

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(docs) == 0 {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	doc := docs[0]

	_, err = doc.Ref.Delete(ctx)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "User deleted successfully"})

}
