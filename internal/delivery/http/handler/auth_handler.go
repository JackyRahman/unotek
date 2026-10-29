package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"unotek-backend/internal/delivery/http/dto"
	"unotek-backend/internal/delivery/http/middleware"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/response"
)

type AuthHandler struct {
	authUsecase domain.AuthUsecase
}

func NewAuthHandler(authUsecase domain.AuthUsecase) *AuthHandler {
	return &AuthHandler{authUsecase: authUsecase}
}

// Login godoc
// @Summary      Authenticate user and obtain JWT token
// @Description  Login with username/email and password to receive access token
// @Tags         Authentication
// @Accept       json
// @Produce      json
// @Param        request body dto.LoginRequest true "Login Credentials"
// @Success      200 {object} response.APIResponse{data=domain.AuthResponse} "Login Successful"
// @Failure      400 {object} response.APIResponse "Bad Request / Validation Error"
// @Failure      401 {object} response.APIResponse "Invalid Credentials"
// @Failure      403 {object} response.APIResponse "User Account Inactive"
// @Router       /api/v1/auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	authResp, err := h.authUsecase.Login(c.Request.Context(), req.Identifier, req.Password)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Login successful", authResp)
}

// GetProfile godoc
// @Summary      Get current logged-in user profile
// @Description  Returns user details, assigned roles, and permission list
// @Tags         Authentication
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=domain.AuthUser} "Profile Retrieved"
// @Failure      401 {object} response.APIResponse "Unauthorized"
// @Router       /api/v1/auth/profile [get]
func (h *AuthHandler) GetProfile(c *gin.Context) {
	userIDVal, exists := c.Get(middleware.CtxUserIDKey)
	if !exists {
		response.Error(c, appErrors.NewUnauthorizedError("Unauthorized"))
		return
	}

	userID, ok := userIDVal.(uuid.UUID)
	if !ok {
		response.Error(c, appErrors.NewUnauthorizedError("Invalid user session"))
		return
	}

	profile, err := h.authUsecase.GetProfile(c.Request.Context(), userID)
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Profile retrieved successfully", profile)
}
