package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go/go-goinc/service"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		const AUTH_SCHEMA = "Bearer "
		authHeader := ctx.GetHeader("Authorization")

		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Not Authorized",
			})
		}

		jwtToken := authHeader[len(AUTH_SCHEMA):]

		jwtService := service.NewJWTService()

		token, err := jwtService.ValidateToken(jwtToken)

		if err != nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		} else {
			claims := token.Claims.(jwt.MapClaims)
			fmt.Println(claims.GetExpirationTime())
			fmt.Println(claims.GetIssuedAt())
			fmt.Println(claims.GetNotBefore())
			fmt.Println(claims.GetIssuer())
			fmt.Println(claims["Name"])
			fmt.Println(claims["Admin"])
		}

	}
}
