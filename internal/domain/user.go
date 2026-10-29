package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/pkg/pagination"
)

type User struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Email        string         `gorm:"type:varchar(255);uniqueIndex;not null" json:"email"`
	Username     string         `gorm:"type:varchar(100);uniqueIndex;not null" json:"username"`
	PasswordHash string         `gorm:"type:varchar(255);not null" json:"-"`
	FullName     string         `gorm:"type:varchar(150);not null" json:"full_name"`
	Phone        string         `gorm:"type:varchar(50)" json:"phone"`
	AvatarURL    string         `gorm:"type:text" json:"avatar_url"`
	IsActive     bool           `gorm:"default:true;index" json:"is_active"`
	LastLoginAt  *time.Time     `json:"last_login_at"`
	Roles        []Role         `gorm:"many2many:user_roles;" json:"roles,omitempty"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type UserRole struct {
	UserID uuid.UUID `gorm:"type:uuid;primaryKey" json:"user_id"`
	RoleID uuid.UUID `gorm:"type:uuid;primaryKey" json:"role_id"`
}

type UserRepository interface {
	GetByID(ctx context.Context, id uuid.UUID) (*User, error)
	GetByEmail(ctx context.Context, email string) (*User, error)
	GetByUsername(ctx context.Context, username string) (*User, error)
	GetByUsernameOrEmail(ctx context.Context, identifier string) (*User, error)
	List(ctx context.Context, params pagination.Params) ([]User, int64, error)
	Create(ctx context.Context, user *User) error
	Update(ctx context.Context, user *User) error
	Delete(ctx context.Context, id uuid.UUID) error
	UpdateLastLogin(ctx context.Context, id uuid.UUID) error
	AssignRoles(ctx context.Context, userID uuid.UUID, roleIDs []uuid.UUID) error
}
