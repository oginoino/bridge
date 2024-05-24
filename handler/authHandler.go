package handler

import (
	"context"
	"net/http"

	"github.com/GinoCodeSpace/bridge/models"
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

		c.Set("email", token.Claims["email"])

		c.Set("name", token.Claims["name"])

		c.Next()
	}
}

func CheckAdminFunction(ctx context.Context, c *gin.Context) bool {
	uid, _ := c.Get("uid")
	var admin models.AdminUser

	documentUserSnapShot, err := dbClient.Collection("admin").Where("uid", "==", uid).Documents(ctx).Next()

	if err != nil {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return true
	}

	documentUserSnapShot.DataTo(&admin)

	if !admin.IsActivated {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return true
	}

	if admin.Role != "admin" {
		sendError(c, http.StatusUnauthorized, "You are not authorized to create a product")
		return true
	}
	return false
}
