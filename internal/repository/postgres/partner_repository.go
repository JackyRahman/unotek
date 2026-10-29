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

type partnerRepository struct {
	db *gorm.DB
}

func NewPartnerRepository(db *gorm.DB) domain.PartnerRepository {
	return &partnerRepository{db: db}
}

func (r *partnerRepository) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.ProductPartner, int64, error) {
	var partners []domain.ProductPartner
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.ProductPartner{})

	if publicOnly {
		query = query.Where("is_active = ?", true)
	}

	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("name ILIKE ? OR partner_type ILIKE ? OR description ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := params.SortBy + " " + params.Order
	if publicOnly {
		orderClause = "display_order ASC, name ASC"
	}

	err := query.Order(orderClause).
		Offset(params.Offset()).
		Limit(params.PerPage).
		Find(&partners).Error

	return partners, total, err
}

func (r *partnerRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProductPartner, error) {
	var partner domain.ProductPartner
	err := r.db.WithContext(ctx).First(&partner, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Product partner not found")
		}
		return nil, err
	}
	return &partner, nil
}

func (r *partnerRepository) GetBySlug(ctx context.Context, slug string) (*domain.ProductPartner, error) {
	var partner domain.ProductPartner
	err := r.db.WithContext(ctx).First(&partner, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Product partner not found")
		}
		return nil, err
	}
	return &partner, nil
}

func (r *partnerRepository) Create(ctx context.Context, partner *domain.ProductPartner) error {
	return r.db.WithContext(ctx).Create(partner).Error
}

func (r *partnerRepository) Update(ctx context.Context, partner *domain.ProductPartner) error {
	return r.db.WithContext(ctx).Save(partner).Error
}

func (r *partnerRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.ProductPartner{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Product partner not found")
	}
	return nil
}
