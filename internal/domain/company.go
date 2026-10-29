package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompanyProfile struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CompanyName     string    `gorm:"type:varchar(150);not null" json:"company_name"`
	Tagline         string    `gorm:"type:varchar(255)" json:"tagline"`
	AboutText       string    `gorm:"type:text" json:"about_text"`
	Vision          string    `gorm:"type:text" json:"vision"`
	Mission         string    `gorm:"type:text" json:"mission"`
	FoundedYear     int       `json:"founded_year"`
	LogoURL         string    `gorm:"type:text" json:"logo_url"`
	FaviconURL      string    `gorm:"type:text" json:"favicon_url"`
	MetaTitle       string    `gorm:"type:varchar(255)" json:"meta_title"`
	MetaDescription string    `gorm:"type:text" json:"meta_description"`
	CreatedAt       time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt       time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type OfficeLocation struct {
	ID            uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	BranchName    string         `gorm:"type:varchar(100);not null" json:"branch_name"`
	BranchType    string         `gorm:"type:varchar(50);default:'head_office'" json:"branch_type"`
	Address       string         `gorm:"type:text;not null" json:"address"`
	City          string         `gorm:"type:varchar(100);not null" json:"city"`
	Province      string         `gorm:"type:varchar(100)" json:"province"`
	PostalCode    string         `gorm:"type:varchar(20)" json:"postal_code"`
	Country       string         `gorm:"type:varchar(100);default:'Indonesia'" json:"country"`
	Phone         string         `gorm:"type:varchar(50)" json:"phone"`
	Email         string         `gorm:"type:varchar(150)" json:"email"`
	WorkingHours  string         `gorm:"type:varchar(100)" json:"working_hours"`
	GoogleMapsURL string         `gorm:"type:text" json:"google_maps_url"`
	Latitude      *float64       `gorm:"type:decimal(10,8)" json:"latitude"`
	Longitude     *float64       `gorm:"type:decimal(11,8)" json:"longitude"`
	IsHeadOffice  bool           `gorm:"default:false" json:"is_head_office"`
	DisplayOrder  int            `gorm:"default:0;index" json:"display_order"`
	IsActive      bool           `gorm:"default:true;index" json:"is_active"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

type CompanyRepository interface {
	GetProfile(ctx context.Context) (*CompanyProfile, error)
	UpdateProfile(ctx context.Context, profile *CompanyProfile) error

	ListOffices(ctx context.Context, publicOnly bool) ([]OfficeLocation, error)
	GetOfficeByID(ctx context.Context, id uuid.UUID) (*OfficeLocation, error)
	CreateOffice(ctx context.Context, office *OfficeLocation) error
	UpdateOffice(ctx context.Context, office *OfficeLocation) error
	DeleteOffice(ctx context.Context, id uuid.UUID) error
}

type CompanyUsecase interface {
	GetProfile(ctx context.Context) (*CompanyProfile, error)
	UpdateProfile(ctx context.Context, profile *CompanyProfile) error

	ListOffices(ctx context.Context, publicOnly bool) ([]OfficeLocation, error)
	GetOfficeByID(ctx context.Context, id uuid.UUID) (*OfficeLocation, error)
	CreateOffice(ctx context.Context, office *OfficeLocation) error
	UpdateOffice(ctx context.Context, office *OfficeLocation) error
	DeleteOffice(ctx context.Context, id uuid.UUID) error
}
