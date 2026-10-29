package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"unotek-backend/pkg/response"
)

func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("[PANIC RECOVERED] %v\n", r)
				c.JSON(http.StatusInternalServerError, response.APIResponse{
					Success: false,
					Message: "An unexpected system error occurred. Please try again later.",
				})
				c.Abort()
			}
		}()
		c.Next()
	}
}
