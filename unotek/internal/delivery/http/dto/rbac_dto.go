package dto

import "github.com/google/uuid"

type CreateRoleRequest struct {
	Name        string `json:"name" binding:"required" example:"marketing"`
	DisplayName string `json:"display_name" binding:"required" example:"Marketing Specialist"`
	Description string `json:"description" example:"Role for marketing content updates"`
}

type UpdateRoleRequest struct {
	DisplayName string `json:"display_name" binding:"required" example:"Marketing Specialist"`
	Description string `json:"description" example:"Role for marketing content updates"`
}

type AssignPermissionsRequest struct {
	PermissionIDs []uuid.UUID `json:"permission_ids" binding:"required"`
}
