package middlewares

import (
	"amarhrs/ecommerce/helpers"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("my-new-super-secret-key")

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("Authorization")

		if tokenString == "" {
			helpers.Error(ctx, http.StatusUnauthorized, "Missing token")
			ctx.Abort()
			return
		}

		// Kalau formatnya "Bearer token", ambil tokennya saja
		if strings.HasPrefix(strings.ToLower(tokenString), "bearer ") {
			tokenString = strings.TrimSpace(tokenString[7:])
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			helpers.Error(ctx, http.StatusUnauthorized, "Invalid token")
			ctx.Abort()
			return
		}

		// Jika token valid, ambil informasi user
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["user_id"].(float64))
		ctx.Set("user_id", userID)

		ctx.Next()
	}
}
