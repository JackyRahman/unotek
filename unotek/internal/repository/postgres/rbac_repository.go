package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
)

type rbacRepository struct {
	db *gorm.DB
}

func NewRBACRepository(db *gorm.DB) domain.RBACRepository {
	return &rbacRepository{db: db}
}

func (r *rbacRepository) GetAllRoles(ctx context.Context) ([]domain.Role, error) {
	var roles []domain.Role
	err := r.db.WithContext(ctx).Preload("Permissions").Find(&roles).Error
	return roles, err
}

func (r *rbacRepository) GetRoleByID(ctx context.Context, id uuid.UUID) (*domain.Role, error) {
	var role domain.Role
	err := r.db.WithContext(ctx).Preload("Permissions").First(&role, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *rbacRepository) GetRoleByName(ctx context.Context, name string) (*domain.Role, error) {
	var role domain.Role
	err := r.db.WithContext(ctx).Preload("Permissions").First(&role, "name = ?", name).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Role not found")
		}
		return nil, err
	}
	return &role, nil
}

func (r *rbacRepository) CreateRole(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Create(role).Error
}

func (r *rbacRepository) UpdateRole(ctx context.Context, role *domain.Role) error {
	return r.db.WithContext(ctx).Save(role).Error
}

func (r *rbacRepository) DeleteRole(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Role{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Role not found")
	}
	return nil
}

func (r *rbacRepository) GetAllPermissions(ctx context.Context) ([]domain.Permission, error) {
	var permissions []domain.Permission
	err := r.db.WithContext(ctx).Find(&permissions).Error
	return permissions, err
}

func (r *rbacRepository) AssignPermissionsToRole(ctx context.Context, roleID uuid.UUID, permissionIDs []uuid.UUID) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// Delete existing permissions for role
		if err := tx.Exec("DELETE FROM role_permissions WHERE role_id = ?", roleID).Error; err != nil {
			return err
		}

		if len(permissionIDs) == 0 {
			return nil
		}

		var rolePerms []domain.RolePermission
		for _, permID := range permissionIDs {
			rolePerms = append(rolePerms, domain.RolePermission{
				RoleID:       roleID,
				PermissionID: permID,
			})
		}
		return tx.Table("role_permissions").Create(&rolePerms).Error
	})
}

func (r *rbacRepository) GetPermissionsByRoleNames(ctx context.Context, roleNames []string) ([]string, error) {
	if len(roleNames) == 0 {
		return []string{}, nil
	}

	var permissionNames []string
	err := r.db.WithContext(ctx).Table("permissions").
		Joins("JOIN role_permissions ON role_permissions.permission_id = permissions.id").
		Joins("JOIN roles ON roles.id = role_permissions.role_id").
		Where("roles.name IN ?", roleNames).
		Distinct("permissions.name").
		Pluck("permissions.name", &permissionNames).Error

	return permissionNames, err
}
