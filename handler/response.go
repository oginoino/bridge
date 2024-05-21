package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, ErrorResponse{
		Message:   msg,
		ErrorCode: fmt.Sprintf("%d", code),
	})
}

func sendSuccess(ctx *gin.Context, message string, code int, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, SuccessResponse{
		Message: message,
		Data:    data,
		Code:    code,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}
