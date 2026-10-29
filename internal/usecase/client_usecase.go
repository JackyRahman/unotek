package usecase

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
)

type clientUsecase struct {
	clientRepo domain.ClientRepository
}

func NewClientUsecase(clientRepo domain.ClientRepository) domain.ClientUsecase {
	return &clientUsecase{clientRepo: clientRepo}
}

func (u *clientUsecase) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.Client, int64, error) {
	return u.clientRepo.List(ctx, params, publicOnly)
}

func (u *clientUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Client, error) {
	return u.clientRepo.GetByID(ctx, id)
}

func (u *clientUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Client, error) {
	return u.clientRepo.GetBySlug(ctx, slug)
}

func (u *clientUsecase) Create(ctx context.Context, client *domain.Client) error {
	slug := sanitizeSlug(client.Slug)
	if slug == "" {
		slug = generateSlug(client.Name)
	}
	client.Slug = slug

	existing, _ := u.clientRepo.GetBySlug(ctx, client.Slug)
	if existing != nil {
		return appErrors.NewConflictError("Client with this slug already exists")
	}

	return u.clientRepo.Create(ctx, client)
}

func (u *clientUsecase) Update(ctx context.Context, client *domain.Client) error {
	existing, err := u.clientRepo.GetByID(ctx, client.ID)
	if err != nil {
		return err
	}

	if client.Slug != "" && client.Slug != existing.Slug {
		existingSlug, _ := u.clientRepo.GetBySlug(ctx, client.Slug)
		if existingSlug != nil && existingSlug.ID != client.ID {
			return appErrors.NewConflictError("Client with this slug already exists")
		}
		existing.Slug = sanitizeSlug(client.Slug)
	}

	existing.Name = client.Name
	existing.Industry = client.Industry
	existing.WebsiteURL = client.WebsiteURL
	existing.LogoURL = client.LogoURL
	existing.Description = client.Description
	existing.IsFeatured = client.IsFeatured
	existing.DisplayOrder = client.DisplayOrder
	existing.IsActive = client.IsActive

	return u.clientRepo.Update(ctx, existing)
}

func (u *clientUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.clientRepo.Delete(ctx, id)
}

func sanitizeSlug(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.ReplaceAll(s, " ", "-")
	return s
}

func generateSlug(s string) string {
	return sanitizeSlug(s)
}
