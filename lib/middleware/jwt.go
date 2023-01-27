package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"walpapperCollectRestAPI/config"
)

func validateAccessJWT(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}
	return config.JWT_KEY, nil
}
func JWT(context *gin.Context) {
	tokenString := context.Request.Header.Get("Cookie")
	if tokenString == "" {
		context.JSON(401, gin.H{"error": "request does not contain an access token"})
		context.Abort()
		return
	}

	vals := strings.Split(tokenString, "=")

	token, err := jwt.ParseWithClaims(vals[1], &config.Claims{}, validateAccessJWT)

	if err != nil {
		context.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		context.Set("id", claims.Id)
		context.Set("user_name", claims.UserName)
	}
	context.Next()
}
