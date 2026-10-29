package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
)

type teamRepository struct {
	db *gorm.DB
}

func NewTeamRepository(db *gorm.DB) domain.TeamRepository {
	return &teamRepository{db: db}
}

func (r *teamRepository) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.TeamMember, int64, error) {
	var members []domain.TeamMember
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.TeamMember{})

	if publicOnly {
		query = query.Where("is_active = ?", true)
	}

	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("full_name ILIKE ? OR position ILIKE ? OR department ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := params.SortBy + " " + params.Order
	if publicOnly {
		orderClause = "display_order ASC, full_name ASC"
	}

	err := query.Order(orderClause).
		Offset(params.Offset()).
		Limit(params.PerPage).
		Find(&members).Error

	return members, total, err
}

func (r *teamRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.TeamMember, error) {
	var member domain.TeamMember
	err := r.db.WithContext(ctx).First(&member, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Team member not found")
		}
		return nil, err
	}
	return &member, nil
}

func (r *teamRepository) Create(ctx context.Context, member *domain.TeamMember) error {
	return r.db.WithContext(ctx).Create(member).Error
}

func (r *teamRepository) Update(ctx context.Context, member *domain.TeamMember) error {
	return r.db.WithContext(ctx).Save(member).Error
}

func (r *teamRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.TeamMember{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Team member not found")
	}
	return nil
}
