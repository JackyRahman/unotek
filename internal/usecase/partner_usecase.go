package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
)

type partnerUsecase struct {
	partnerRepo domain.PartnerRepository
}

func NewPartnerUsecase(partnerRepo domain.PartnerRepository) domain.PartnerUsecase {
	return &partnerUsecase{partnerRepo: partnerRepo}
}

func (u *partnerUsecase) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.ProductPartner, int64, error) {
	return u.partnerRepo.List(ctx, params, publicOnly)
}

func (u *partnerUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.ProductPartner, error) {
	return u.partnerRepo.GetByID(ctx, id)
}

func (u *partnerUsecase) GetBySlug(ctx context.Context, slug string) (*domain.ProductPartner, error) {
	return u.partnerRepo.GetBySlug(ctx, slug)
}

func (u *partnerUsecase) Create(ctx context.Context, partner *domain.ProductPartner) error {
	slug := sanitizeSlug(partner.Slug)
	if slug == "" {
		slug = generateSlug(partner.Name)
	}
	partner.Slug = slug

	existing, _ := u.partnerRepo.GetBySlug(ctx, partner.Slug)
	if existing != nil {
		return appErrors.NewConflictError("Product partner with this slug already exists")
	}

	return u.partnerRepo.Create(ctx, partner)
}

func (u *partnerUsecase) Update(ctx context.Context, partner *domain.ProductPartner) error {
	existing, err := u.partnerRepo.GetByID(ctx, partner.ID)
	if err != nil {
		return err
	}

	if partner.Slug != "" && partner.Slug != existing.Slug {
		existingSlug, _ := u.partnerRepo.GetBySlug(ctx, partner.Slug)
		if existingSlug != nil && existingSlug.ID != partner.ID {
			return appErrors.NewConflictError("Product partner with this slug already exists")
		}
		existing.Slug = sanitizeSlug(partner.Slug)
	}

	existing.Name = partner.Name
	existing.PartnerType = partner.PartnerType
	existing.WebsiteURL = partner.WebsiteURL
	existing.LogoURL = partner.LogoURL
	existing.Description = partner.Description
	existing.DisplayOrder = partner.DisplayOrder
	existing.IsActive = partner.IsActive

	return u.partnerRepo.Update(ctx, existing)
}

func (u *partnerUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.partnerRepo.Delete(ctx, id)
}
