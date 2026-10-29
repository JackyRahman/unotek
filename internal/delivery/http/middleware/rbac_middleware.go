package middleware

import (
	"github.com/gin-gonic/gin"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/response"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRolesVal, exists := c.Get(CtxUserRolesKey)
		if !exists {
			response.Error(c, appErrors.NewForbiddenError("Access denied: No roles found"))
			c.Abort()
			return
		}

		userRoles, ok := userRolesVal.([]string)
		if !ok {
			response.Error(c, appErrors.NewForbiddenError("Access denied: Invalid roles format"))
			c.Abort()
			return
		}

		for _, userRole := range userRoles {
			if userRole == "super_admin" {
				c.Next()
				return
			}
			for _, allowedRole := range allowedRoles {
				if userRole == allowedRole {
					c.Next()
					return
				}
			}
		}

		response.Error(c, appErrors.NewForbiddenError("Access denied: Insufficient role permissions"))
		c.Abort()
	}
}

func RequirePermission(requiredPermission string) gin.HandlerFunc {
	return func(c *gin.Context) {
		userRolesVal, exists := c.Get(CtxUserRolesKey)
		if exists {
			if userRoles, ok := userRolesVal.([]string); ok {
				for _, r := range userRoles {
					if r == "super_admin" {
						c.Next()
						return
					}
				}
			}
		}

		userPermsVal, exists := c.Get(CtxUserPermissionsKey)
		if !exists {
			response.Error(c, appErrors.NewForbiddenError("Access denied: No permissions found"))
			c.Abort()
			return
		}

		userPermissions, ok := userPermsVal.([]string)
		if !ok {
			response.Error(c, appErrors.NewForbiddenError("Access denied: Invalid permissions format"))
			c.Abort()
			return
		}

		for _, perm := range userPermissions {
			if perm == requiredPermission {
				c.Next()
				return
			}
		}

		response.Error(c, appErrors.NewForbiddenError("Access denied: Missing required permission '"+requiredPermission+"'"))
		c.Abort()
	}
}
