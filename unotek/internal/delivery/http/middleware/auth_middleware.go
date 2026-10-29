package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/response"
	"unotek-backend/pkg/security"
)

const (
	CtxUserIDKey         = "user_id"
	CtxUserEmailKey      = "user_email"
	CtxUsernameKey       = "username"
	CtxUserRolesKey      = "user_roles"
	CtxUserPermissionsKey = "user_permissions"
)

func AuthMiddleware(jwtService security.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, appErrors.NewUnauthorizedError("Authorization header is required"))
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			response.Error(c, appErrors.NewUnauthorizedError("Invalid Authorization format. Expected 'Bearer <token>'"))
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims, err := jwtService.ValidateToken(tokenString)
		if err != nil {
			response.Error(c, err)
			c.Abort()
			return
		}

		c.Set(CtxUserIDKey, claims.UserID)
		c.Set(CtxUserEmailKey, claims.Email)
		c.Set(CtxUsernameKey, claims.Username)
		c.Set(CtxUserRolesKey, claims.Roles)
		c.Set(CtxUserPermissionsKey, claims.Permissions)

		c.Next()
	}
}
