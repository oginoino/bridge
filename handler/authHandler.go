package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *AuthHandler) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenValue := c.Request.Header.Get("Authorization")

		if tokenValue == "" {
			sendError(c, http.StatusUnauthorized, "Authorization header is missing")
			return
		}

		token, err := handler.authClient.VerifyIDToken(handler.ctx, tokenValue)

		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.Set("token", token)

		c.Set("uid", token.UID)

		c.Next()
	}
}
