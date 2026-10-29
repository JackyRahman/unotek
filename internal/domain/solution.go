package domain

import (
	"context"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/pkg/pagination"
)

type Solution struct {
	ID           uuid.UUID         `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	Title        string            `gorm:"type:varchar(150);not null" json:"title"`
	Slug         string            `gorm:"type:varchar(180);uniqueIndex;not null" json:"slug"`
	Tagline      string            `gorm:"type:varchar(255)" json:"tagline"`
	Summary      string            `gorm:"type:text" json:"summary"`
	Description  string            `gorm:"type:text" json:"description"`
	IconURL      string            `gorm:"type:text" json:"icon_url"`
	HeroImageURL string            `gorm:"type:text" json:"hero_image_url"`
	BrochureURL  string            `gorm:"type:text" json:"brochure_url"`
	DisplayOrder int               `gorm:"default:0;index" json:"display_order"`
	IsFeatured   bool              `gorm:"default:false;index" json:"is_featured"`
	IsActive     bool              `gorm:"default:true;index" json:"is_active"`
	Features     []SolutionFeature `gorm:"foreignKey:SolutionID;constraint:OnDelete:CASCADE" json:"features,omitempty"`
	Modules      []SolutionModule  `gorm:"foreignKey:SolutionID;constraint:OnDelete:CASCADE" json:"modules,omitempty"`
	CreatedAt    time.Time         `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time         `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt    gorm.DeletedAt    `gorm:"index" json:"deleted_at,omitempty"`
}

type SolutionFeature struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SolutionID   uuid.UUID `gorm:"type:uuid;not null;index" json:"solution_id"`
	Title        string    `gorm:"type:varchar(150);not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	IconURL      string    `gorm:"type:text" json:"icon_url"`
	DisplayOrder int       `gorm:"default:0;index" json:"display_order"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SolutionModule struct {
	ID           uuid.UUID `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	SolutionID   uuid.UUID `gorm:"type:uuid;not null;index" json:"solution_id"`
	Title        string    `gorm:"type:varchar(150);not null" json:"title"`
	Description  string    `gorm:"type:text" json:"description"`
	ImageURL     string    `gorm:"type:text" json:"image_url"`
	DisplayOrder int       `gorm:"default:0;index" json:"display_order"`
	IsActive     bool      `gorm:"default:true" json:"is_active"`
	CreatedAt    time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt    time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

type SolutionRepository interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]Solution, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Solution, error)
	GetBySlug(ctx context.Context, slug string) (*Solution, error)
	Create(ctx context.Context, solution *Solution) error
	Update(ctx context.Context, solution *Solution) error
	Delete(ctx context.Context, id uuid.UUID) error

	// Feature operations
	AddFeature(ctx context.Context, feature *SolutionFeature) error
	UpdateFeature(ctx context.Context, feature *SolutionFeature) error
	DeleteFeature(ctx context.Context, id uuid.UUID) error

	// Module operations
	AddModule(ctx context.Context, module *SolutionModule) error
	UpdateModule(ctx context.Context, module *SolutionModule) error
	DeleteModule(ctx context.Context, id uuid.UUID) error
}

type SolutionUsecase interface {
	List(ctx context.Context, params pagination.Params, publicOnly bool) ([]Solution, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*Solution, error)
	GetBySlug(ctx context.Context, slug string) (*Solution, error)
	Create(ctx context.Context, solution *Solution) error
	Update(ctx context.Context, solution *Solution) error
	Delete(ctx context.Context, id uuid.UUID) error

	AddFeature(ctx context.Context, feature *SolutionFeature) error
	UpdateFeature(ctx context.Context, feature *SolutionFeature) error
	DeleteFeature(ctx context.Context, id uuid.UUID) error

	AddModule(ctx context.Context, module *SolutionModule) error
	UpdateModule(ctx context.Context, module *SolutionModule) error
	DeleteModule(ctx context.Context, id uuid.UUID) error
}
