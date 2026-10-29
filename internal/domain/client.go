package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/pkg/pagination"
)

type Client struct {
	ID           uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Name         string         `gorm:"type:varchar(150);not null" json:"name"`
	Slug         string         `gorm:"type:varchar(180);uniqueIndex;not null" json:"slug"`
	Industry     string         `gorm:"type:varchar(100)" json:"industry"`
	WebsiteURL   string         `gorm:"type:text" json:"website_url"`
	LogoURL      string         `gorm:"type:text;not null" json:"logo_url"`
	Description  string         `gorm:"type:text" json:"description"`
	IsFeatured   bool           `gorm:"default:false;index" json:"is_featured"`
	DisplayOrder int            `gorm:"default:0;index" json:"display_order"`
	IsActive     bool           `gorm:"default:true;index" json:"is_active"`
	CreatedAt    time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type ClientRepository interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]Client, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Client, error)
	GetBySlug(ctx context.Context, slug string) (*Client, error)
	Create(ctx context.Context, client *Client) error
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ClientUsecase interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]Client, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Client, error)
	GetBySlug(ctx context.Context, slug string) (*Client, error)
	Create(ctx context.Context, client *Client) error
	Update(ctx context.Context, client *Client) error
	Delete(ctx context.Context, id uuid.UUID) error
}
