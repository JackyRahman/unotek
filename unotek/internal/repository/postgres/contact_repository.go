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

type contactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) domain.ContactRepository {
	return &contactRepository{db: db}
}

func (r *contactRepository) List(ctx context.Context, params pagination.Params, statusFilter string) ([]domain.ContactSubmission, int64, error) {
	var submissions []domain.ContactSubmission
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.ContactSubmission{})

	if statusFilter != "" {
		query = query.Where("status = ?", statusFilter)
	}

	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("full_name ILIKE ? OR email ILIKE ? OR company_name ILIKE ? OR subject ILIKE ? OR message ILIKE ?", searchPattern, searchPattern, searchPattern, searchPattern, searchPattern)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	orderClause := params.SortBy + " " + params.Order
	err := query.Order(orderClause).
		Offset(params.Offset()).
		Limit(params.PerPage).
		Find(&submissions).Error

	return submissions, total, err
}

func (r *contactRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.ContactSubmission, error) {
	var submission domain.ContactSubmission
	err := r.db.WithContext(ctx).First(&submission, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Contact submission not found")
		}
		return nil, err
	}
	return &submission, nil
}

func (r *contactRepository) Create(ctx context.Context, contact *domain.ContactSubmission) error {
	return r.db.WithContext(ctx).Create(contact).Error
}

func (r *contactRepository) UpdateStatus(ctx context.Context, id uuid.UUID, status, notes string) error {
	updates := map[string]interface{}{
		"status": status,
	}
	if notes != "" {
		updates["notes"] = notes
	}

	result := r.db.WithContext(ctx).Model(&domain.ContactSubmission{}).
		Where("id = ?", id).
		Updates(updates)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Contact submission not found")
	}
	return nil
}

func (r *contactRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.ContactSubmission{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Contact submission not found")
	}
	return nil
}
