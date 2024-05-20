package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.Request.Header.Get("Authorization")

		if tokenValue == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "No token provided"})
			return
		}

		token, err := handler.authClient.VerifyIDToken(handler.ctx, tokenValue)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		fmt.Println("Usuário autenticado " + token.UID)

		c.Next()
	}
}
