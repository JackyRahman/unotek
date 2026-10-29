package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
)

type companyUsecase struct {
	companyRepo domain.CompanyRepository
}

func NewCompanyUsecase(companyRepo domain.CompanyRepository) domain.CompanyUsecase {
	return &companyUsecase{companyRepo: companyRepo}
}

func (u *companyUsecase) GetProfile(ctx context.Context) (*domain.CompanyProfile, error) {
	return u.companyRepo.GetProfile(ctx)
}

func (u *companyUsecase) UpdateProfile(ctx context.Context, profile *domain.CompanyProfile) error {
	return u.companyRepo.UpdateProfile(ctx, profile)
}

func (u *companyUsecase) ListOffices(ctx context.Context, publicOnly bool) ([]domain.OfficeLocation, error) {
	return u.companyRepo.ListOffices(ctx, publicOnly)
}

func (u *companyUsecase) GetOfficeByID(ctx context.Context, id uuid.UUID) (*domain.OfficeLocation, error) {
	return u.companyRepo.GetOfficeByID(ctx, id)
}

func (u *companyUsecase) CreateOffice(ctx context.Context, office *domain.OfficeLocation) error {
	return u.companyRepo.CreateOffice(ctx, office)
}

func (u *companyUsecase) UpdateOffice(ctx context.Context, office *domain.OfficeLocation) error {
	existing, err := u.companyRepo.GetOfficeByID(ctx, office.ID)
	if err != nil {
		return err
	}

	existing.BranchName = office.BranchName
	existing.BranchType = office.BranchType
	existing.Address = office.Address
	existing.City = office.City
	existing.Province = office.Province
	existing.PostalCode = office.PostalCode
	existing.Country = office.Country
	existing.Phone = office.Phone
	existing.Email = office.Email
	existing.WorkingHours = office.WorkingHours
	existing.GoogleMapsURL = office.GoogleMapsURL
	existing.Latitude = office.Latitude
	existing.Longitude = office.Longitude
	existing.IsHeadOffice = office.IsHeadOffice
	existing.DisplayOrder = office.DisplayOrder
	existing.IsActive = office.IsActive

	return u.companyRepo.UpdateOffice(ctx, existing)
}

func (u *companyUsecase) DeleteOffice(ctx context.Context, id uuid.UUID) error {
	return u.companyRepo.DeleteOffice(ctx, id)
}
