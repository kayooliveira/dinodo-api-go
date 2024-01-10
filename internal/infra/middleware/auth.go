package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kayooliveira/dinodo-api-go/internal/infra/auth"
)

func AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader("Authorization")

	if len(authHeader) <= 0 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
		})

		ctx.Abort()
		return
	}

	token := strings.Split(authHeader, "Bearer ")[1]

	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
			"message":   "Please provide a Bearer token",
		})
		ctx.Abort()
		return
	}

	err := auth.ParseUserToken(token)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
			"message":   err.Error(),
		})

		ctx.Abort()
	}

	ctx.Next()
}
