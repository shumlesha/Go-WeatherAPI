package middleware

import (
	"WeatherfForecast/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	bearerPrefix        = "Bearer"
	userId              = "userId"
	username            = "username"
)

type AuthMiddleware struct {
	jwtUtil util.JwtUtil
}

func NewAuthMiddleware(jwtUtil util.JwtUtil) *AuthMiddleware {
	return &AuthMiddleware{
		jwtUtil: jwtUtil,
	}
}

func (m *AuthMiddleware) AuthChecked() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader(authorizationHeader)

		if authHeader == "" {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				util.NewErrorResponse("Unauthorized", "Authorization header is missing"))
			return
		}

		headerParts := strings.SplitN(authHeader, " ", 2)
		if !(len(headerParts) == 2 && headerParts[0] == bearerPrefix) {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				util.NewErrorResponse("Unauthorized", "Invalid authorization header"))
			return
		}

		claims, err := m.jwtUtil.ValidateToken(headerParts[1])

		if err != nil {
			context.AbortWithStatusJSON(
				http.StatusUnauthorized,
				util.NewErrorResponse("Unauthorized", "Invalid token"))
			return
		}

		context.Set(userId, claims.UserId)
		context.Set(username, claims.Username)

		context.Next()
	}
}
