package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func SuccessWithPagination(c *gin.Context, statusCode int, message string, data interface{}, meta *pagination.Meta) {
	c.JSON(statusCode, APIResponse{
		Success: true,
		Message: message,
		Data:    data,
		Meta:    meta,
	})
}

func Error(c *gin.Context, err error) {
	if appErr, ok := err.(*appErrors.AppError); ok {
		c.JSON(appErr.StatusCode, APIResponse{
			Success: false,
			Message: appErr.Message,
			Error:   appErr.Error(),
		})
		return
	}

	c.JSON(http.StatusInternalServerError, APIResponse{
		Success: false,
		Message: "Internal server error",
		Error:   err.Error(),
	})
}

func ValidationError(c *gin.Context, message string, details interface{}) {
	c.JSON(http.StatusUnprocessableEntity, APIResponse{
		Success: false,
		Message: message,
		Error:   details,
	})
}
