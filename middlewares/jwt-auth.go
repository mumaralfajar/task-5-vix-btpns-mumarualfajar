package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"task-5-vix-btpns-mumaralfajar/services"
)

func AuthMiddleware(jwtService services.JwtService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			response := gin.H{"error": "unauthorized"}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		token, err := jwtService.ValidateToken(authHeader)
		if err != nil {
			response := gin.H{"error": err.Error()}
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims := token.Claims.(services.JwtCustomClaims)
		ctx.Set("currentUser", claims.ID)
		ctx.Next()
	}
}
