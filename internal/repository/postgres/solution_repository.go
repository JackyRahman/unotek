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

type solutionRepository struct {
	db *gorm.DB
}

func NewSolutionRepository(db *gorm.DB) domain.SolutionRepository {
	return &solutionRepository{db: db}
}

func (r *solutionRepository) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.Solution, int64, error) {
	var solutions []domain.Solution
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Solution{})

	if publicOnly {
		query = query.Where("is_active = ?", true)
	}

	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("title ILIKE ? OR tagline ILIKE ? OR summary ILIKE ?", searchPattern, searchPattern, searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := params.SortBy + " " + params.Order
	if publicOnly {
		orderClause = "display_order ASC, title ASC"
	}

	err := query.Preload("Features", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("display_order ASC")
	}).Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("display_order ASC")
	}).Order(orderClause).
		Offset(params.Offset()).
		Limit(params.PerPage).
		Find(&solutions).Error

	return solutions, total, err
}

func (r *solutionRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Solution, error) {
	var solution domain.Solution
	err := r.db.WithContext(ctx).Preload("Features", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Order("display_order ASC")
	}).First(&solution, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Solution not found")
		}
		return nil, err
	}
	return &solution, nil
}

func (r *solutionRepository) GetBySlug(ctx context.Context, slug string) (*domain.Solution, error) {
	var solution domain.Solution
	err := r.db.WithContext(ctx).Preload("Features", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("display_order ASC")
	}).Preload("Modules", func(db *gorm.DB) *gorm.DB {
		return db.Where("is_active = ?", true).Order("display_order ASC")
	}).First(&solution, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Solution not found")
		}
		return nil, err
	}
	return &solution, nil
}

func (r *solutionRepository) Create(ctx context.Context, solution *domain.Solution) error {
	return r.db.WithContext(ctx).Create(solution).Error
}

func (r *solutionRepository) Update(ctx context.Context, solution *domain.Solution) error {
	return r.db.WithContext(ctx).Save(solution).Error
}

func (r *solutionRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Solution{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Solution not found")
	}
	return nil
}

func (r *solutionRepository) AddFeature(ctx context.Context, feature *domain.SolutionFeature) error {
	return r.db.WithContext(ctx).Create(feature).Error
}

func (r *solutionRepository) UpdateFeature(ctx context.Context, feature *domain.SolutionFeature) error {
	return r.db.WithContext(ctx).Save(feature).Error
}

func (r *solutionRepository) DeleteFeature(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.SolutionFeature{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Solution feature not found")
	}
	return nil
}

func (r *solutionRepository) AddModule(ctx context.Context, module *domain.SolutionModule) error {
	return r.db.WithContext(ctx).Create(module).Error
}

func (r *solutionRepository) UpdateModule(ctx context.Context, module *domain.SolutionModule) error {
	return r.db.WithContext(ctx).Save(module).Error
}

func (r *solutionRepository) DeleteModule(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.SolutionModule{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Solution module not found")
	}
	return nil
}
