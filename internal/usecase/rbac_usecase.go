package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
)

type RBACUsecase interface {
	GetRoles(ctx context.Context) ([]domain.Role, error)
	GetRoleByID(ctx context.Context, id uuid.UUID) (*domain.Role, error)
	CreateRole(ctx context.Context, role *domain.Role) error
	UpdateRole(ctx context.Context, role *domain.Role) error
	DeleteRole(ctx context.Context, id uuid.UUID) error
	GetPermissions(ctx context.Context) ([]domain.Permission, error)
	AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error
}

type rbacUsecase struct {
	rbacRepo domain.RBACRepository
}

func NewRBACUsecase(rbacRepo domain.RBACRepository) RBACUsecase {
	return &rbacUsecase{rbacRepo: rbacRepo}
}

func (u *rbacUsecase) GetRoles(ctx context.Context) ([]domain.Role, error) {
	return u.rbacRepo.GetAllRoles(ctx)
}

func (u *rbacUsecase) GetRoleByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	return u.rbacRepo.GetRoleByID(ctx, id)
}

func (u *rbacUsecase) CreateRole(ctx context.Context, role *domain.Role) error {
	existing, _ := u.rbacRepo.GetRoleByName(ctx, role.Name)
	if existing != nil {
		return appErrors.NewConflictError("Role with this name already exists")
	}
	return u.rbacRepo.CreateRole(ctx, role)
}

func (u *rbacUsecase) UpdateRole(ctx context.Context, role *domain.Role) error {
	existing, err := u.rbacRepo.GetRoleByID(ctx, role.ID)
	if err != nil {
		return err
	}
	existing.DisplayName = role.DisplayName
	existing.Description = role.Description
	return u.rbacRepo.UpdateRole(ctx, existing)
}

func (u *rbacUsecase) DeleteRole(ctx context.Context, id uuid.UUID) error {
	role, err := u.rbacRepo.GetRoleByID(ctx, id)
	if err != nil {
		return err
	}
	if role.Name == "super_admin" {
		return appErrors.NewForbiddenError("Cannot delete super_admin role")
	}
	return u.rbacRepo.DeleteRole(ctx, id)
}

func (u *rbacUsecase) GetPermissions(ctx context.Context) ([]domain.Permission, error) {
	return u.rbacRepo.GetAllPermissions(ctx)
}

func (u *rbacUsecase) AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	role, err := u.rbacRepo.GetRoleByID(ctx, roleID)
	if err != nil {
		return err
	}
	if role.Name == "super_admin" {
		return appErrors.NewForbiddenError("Super admin role permissions cannot be altered")
	}
	return u.rbacRepo.AssignPermissionsToRole(ctx, roleID, permissionIDs)
}
