package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/pkg/pagination"
)

type TeamMember struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FullName     string         `gorm:"type:varchar(150);not null" json:"full_name"`
	Position     string         `gorm:"type:varchar(100);not null" json:"position"`
	Department   string         `gorm:"type:varchar(100)" json:"department"`
	Bio          string         `gorm:"type:text" json:"bio"`
	PhotoURL     string         `gorm:"type:text" json:"photo_url"`
	LinkedinURL  string         `gorm:"type:text" json:"linkedin_url"`
	Email        string         `gorm:"type:varchar(150)" json:"email"`
	DisplayOrder int            `gorm:"default:0;index" json:"display_order"`
	IsLeadership bool           `gorm:"default:false;index" json:"is_leadership"`
	IsActive     bool           `gorm:"default:true;index" json:"is_active"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type TeamRepository interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]TeamMember, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*TeamMember, error)
	Create(ctx context.Context, member *TeamMember) error
	Update(ctx context.Context, member *TeamMember) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type TeamUsecase interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]TeamMember, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*TeamMember, error)
	Create(ctx context.Context, member *TeamMember) error
	Update(ctx context.Context, member *TeamMember) error
	Delete(ctx context.Context, id uuid.UUID) error
}
