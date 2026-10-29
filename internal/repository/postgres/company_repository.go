package postgres

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
)

type companyRepository struct {
	db *gorm.DB
}

func NewCompanyRepository(db *gorm.DB) domain.CompanyRepository {
	return &companyRepository{db: db}
}

func (r *companyRepository) GetProfile(ctx context.Context) (*domain.CompanyProfile, error) {
	var profile domain.CompanyProfile
	err := r.db.WithContext(ctx).First(&profile).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Company profile not set")
		}
		return nil, err
	}
	return &profile, nil
}

func (r *companyRepository) UpdateProfile(ctx context.Context, profile *domain.CompanyProfile) error {
	var count int64
	r.db.WithContext(ctx).Model(&domain.CompanyProfile{}).Count(&count)
	if count == 0 {
		return r.db.WithContext(ctx).Create(profile).Error
	}
	return r.db.WithContext(ctx).Save(profile).Error
}

func (r *companyRepository) ListOffices(ctx context.Context, publicOnly bool) ([]domain.OfficeLocation, error) {
	var offices []domain.OfficeLocation
	query := r.db.WithContext(ctx)
	if publicOnly {
		query = query.Where("is_active = ?", true)
	}
	err := query.Order("display_order ASC, is_head_office DESC").Find(&offices).Error
	return offices, err
}

func (r *companyRepository) GetOfficeByID(ctx context.Context, id uuid.UUID) (*domain.OfficeLocation, error) {
	var office domain.OfficeLocation
	err := r.db.WithContext(ctx).First(&office, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, appErrors.NewNotFoundError("Office location not found")
		}
		return nil, err
	}
	return &office, nil
}

func (r *companyRepository) CreateOffice(ctx context.Context, office *domain.OfficeLocation) error {
	return r.db.WithContext(ctx).Create(office).Error
}

func (r *companyRepository) UpdateOffice(ctx context.Context, office *domain.OfficeLocation) error {
	return r.db.WithContext(ctx).Save(office).Error
}

func (r *companyRepository) DeleteOffice(ctx context.Context, id uuid.UUID) error {
	result := r.db.WithContext(ctx).Delete(&domain.OfficeLocation{}, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return appErrors.NewNotFoundError("Office location not found")
	}
	return nil
}
