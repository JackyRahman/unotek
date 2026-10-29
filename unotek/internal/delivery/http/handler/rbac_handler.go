package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"unotek-backend/internal/delivery/http/dto"
	"unotek-backend/internal/domain"
	"unotek-backend/internal/usecase"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/response"
)

type RBACHandler struct {
	rbacUsecase usecase.RBACUsecase
}

func NewRBACHandler(rbacUsecase usecase.RBACUsecase) *RBACHandler {
	return &RBACHandler{rbacUsecase: rbacUsecase}
}

// GetRoles godoc
// @Summary      Get all system roles (Super Admin)
// @Description  Retrieve list of roles with associated permissions
// @Tags         Admin - RBAC
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]domain.Role} "List of Roles"
// @Router       /api/v1/admin/rbac/roles [get]
func (h *RBACHandler) GetRoles(c *gin.Context) {
	roles, err := h.rbacUsecase.GetRoles(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Roles retrieved successfully", roles)
}

// CreateRole godoc
// @Summary      Create a new role (Super Admin)
// @Description  Add new custom role to system
// @Tags         Admin - RBAC
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        request body dto.CreateRoleRequest true "Role Data"
// @Success      201 {object} response.APIResponse{data=domain.Role} "Role Created"
// @Router       /api/v1/admin/rbac/roles [post]
func (h *RBACHandler) CreateRole(c *gin.Context) {
	var req dto.CreateRoleRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	role := domain.Role{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		Description: req.Description,
	}

	if err := h.rbacUsecase.CreateRole(c.Request.Context(), &role); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusCreated, "Role created successfully", role)
}

// AssignPermissions godoc
// @Summary      Assign permissions to role (Super Admin)
// @Description  Set list of permissions allowed for a specific role
// @Tags         Admin - RBAC
// @Accept       json
// @Produce      json
// @Security     BearerAuth
// @Param        id path string true "Role UUID"
// @Param        request body dto.AssignPermissionsRequest true "Permission IDs"
// @Success      200 {object} response.APIResponse "Permissions Assigned Successfully"
// @Router       /api/v1/admin/rbac/roles/{id}/permissions [post]
func (h *RBACHandler) AssignPermissions(c *gin.Context) {
	idStr := c.Param("id")
	roleID, err := uuid.Parse(idStr)
	if err != nil {
		response.Error(c, appErrors.NewBadRequestError("Invalid role ID"))
		return
	}

	var req dto.AssignPermissionsRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ValidationError(c, "Invalid input data", err.Error())
		return
	}

	if err := h.rbacUsecase.AssignPermissionsToRole(c.Request.Context(), roleID, req.PermissionIDs); err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Permissions assigned to role successfully", nil)
}

// GetPermissions godoc
// @Summary      Get all available system permissions (Super Admin)
// @Description  List all granular system permissions
// @Tags         Admin - RBAC
// @Produce      json
// @Security     BearerAuth
// @Success      200 {object} response.APIResponse{data=[]domain.Permission} "List of Permissions"
// @Router       /api/v1/admin/rbac/permissions [get]
func (h *RBACHandler) GetPermissions(c *gin.Context) {
	permissions, err := h.rbacUsecase.GetPermissions(c.Request.Context())
	if err != nil {
		response.Error(c, err)
		return
	}

	response.Success(c, http.StatusOK, "Permissions retrieved successfully", permissions)
}
