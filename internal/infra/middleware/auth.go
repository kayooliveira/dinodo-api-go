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

	bearerToken := strings.Split(authHeader, "Bearer ")

	if len(bearerToken) < 2 {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
			"message":   "Please provide a token in format: Bearer token",
		})
		ctx.Abort()
		return
	}

	jwtToken := bearerToken[1]

	if jwtToken == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
			"message":   "Please provide a Bearer token",
		})
		ctx.Abort()
		return
	}

	parsed, err := auth.ParseUserToken(jwtToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error":     "Unauthorized",
			"errorCode": http.StatusUnauthorized,
			"message":   err.Error(),
		})

		ctx.Abort()
	}
	ctx.Set("userId", parsed.ID)
	ctx.Next()
}
