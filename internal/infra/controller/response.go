package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendSuccess(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Success",
		"operation": operation,
		"data":      data,
	})
}

func sendError(ctx *gin.Context, code int, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"error":     "Error",
		"message":   message,
		"errorCode": code,
	})
}

func sendUnauthorized(ctx *gin.Context, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusUnauthorized, gin.H{
		"error":     "Unauthorized",
		"message":   message,
		"errorCode": http.StatusUnauthorized,
	})
}

func sendForbidden(ctx *gin.Context, message string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusForbidden, gin.H{
		"error":     "Forbidden",
		"message":   message,
		"errorCode": http.StatusForbidden,
	})
}

func sendCreated(ctx *gin.Context, operation string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(201, gin.H{
		"message":   "Created",
		"operation": operation,
		"data":      data,
	})
}
