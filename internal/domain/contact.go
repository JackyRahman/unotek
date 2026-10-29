package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/pkg/pagination"
)

type ContactSubmission struct {
	ID              uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	FullName        string         `gorm:"type:varchar(150);not null" json:"full_name"`
	Email           string         `gorm:"type:varchar(150);not null;index" json:"email"`
	Phone           string         `gorm:"type:varchar(50)" json:"phone"`
	CompanyName     string         `gorm:"type:varchar(150)" json:"company_name"`
	ServiceInterest string         `gorm:"type:varchar(100)" json:"service_interest"`
	Subject         string         `gorm:"type:varchar(200);not null" json:"subject"`
	Message         string         `gorm:"type:text;not null" json:"message"`
	Status          string         `gorm:"type:varchar(50);default:'new';index" json:"status"` // new, in_review, contacted, resolved, archived
	Notes           string         `gorm:"type:text" json:"notes"`
	IPAddress       string         `gorm:"type:varchar(45)" json:"ip_address,omitempty"`
	UserAgent       string         `gorm:"type:text" json:"user_agent,omitempty"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type ContactRepository interface {
	List(ctx context.Context, params pagination.Params, statusFilter string) ([]ContactSubmission, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ContactSubmission, error)
	Create(ctx context.Context, contact *ContactSubmission) error
	UpdateStatus(ctx context.Context, id uuid.UUID, status, notes string) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ContactUsecase interface {
	Submit(ctx context.Context, contact *ContactSubmission) error
	List(ctx context.Context, params pagination.Params, statusFilter string) ([]ContactSubmission, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*ContactSubmission, error)
	UpdateStatus(ctx context.Context, id uuid.UUID, status, notes string) error
	Delete(ctx context.Context, id uuid.UUID) error
}
