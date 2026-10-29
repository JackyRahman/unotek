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

type clientRepository struct {
	db *gorm.DB
}

func NewClientRepository(db *gorm.DB) domain.ClientRepository {
	return &clientRepository{db: db}
}

func (r *clientRepository) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.Client, int64, error) {
	var clients []domain.Client
	var total int64

	query := r.db.WithContext(ctx).Model(&domain.Client{})

	if publicOnly {
		query = query.Where("is_active = ?", true)
	}

	if params.Search != "" {
		searchPattern := "%" + params.Search + "%"
		query = query.Where("name ILIKE ? OR industry ILIKE ? OR description ILIKE ?", searchPattern, searchPattern, searchPattern)
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
		Find(&clients).Error

	return clients, total, err
}

func (r *clientRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Client, error) {
	var client domain.Client
	err := r.db.WithContext(ctx).First(&client, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Client not found")
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) GetBySlug(ctx context.Context, slug string) (*domain.Client, error) {
	var client domain.Client
	err := r.db.WithContext(ctx).First(&client, "slug = ?", slug).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Client not found")
		}
		return nil, err
	}
	return &client, nil
}

func (r *clientRepository) Create(ctx context.Context, client *domain.Client) error {
	return r.db.WithContext(ctx).Create(client).Error
}

func (r *clientRepository) Update(ctx context.Context, client *domain.Client) error {
	return r.db.WithContext(ctx).Save(client).Error
}

func (r *clientRepository) Delete(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.Client{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Client not found")
	}
	return nil
}
