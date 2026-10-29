package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/pagination"
)

type solutionUsecase struct {
	solutionRepo domain.SolutionRepository
}

func NewSolutionUsecase(solutionRepo domain.SolutionRepository) domain.SolutionUsecase {
	return &solutionUsecase{solutionRepo: solutionRepo}
}

func (u *solutionUsecase) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.Solution, int64, error) {
	return u.solutionRepo.List(ctx, params, publicOnly)
}

func (u *solutionUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.Solution, error) {
	return u.solutionRepo.GetByID(ctx, id)
}

func (u *solutionUsecase) GetBySlug(ctx context.Context, slug string) (*domain.Solution, error) {
	return u.solutionRepo.GetBySlug(ctx, slug)
}

func (u *solutionUsecase) Create(ctx context.Context, solution *domain.Solution) error {
	slug := sanitizeSlug(solution.Slug)
	if slug == "" {
		slug = generateSlug(solution.Title)
	}
	solution.Slug = slug

	existing, _ := u.solutionRepo.GetBySlug(ctx, solution.Slug)
	if existing != nil {
		return appErrors.NewConflictError("Solution with this slug already exists")
	}

	return u.solutionRepo.Create(ctx, solution)
}

func (u *solutionUsecase) Update(ctx context.Context, solution *domain.Solution) error {
	existing, err := u.solutionRepo.GetByID(ctx, solution.ID)
	if err != nil {
		return err
	}

	if solution.Slug != "" && solution.Slug != existing.Slug {
		existingSlug, _ := u.solutionRepo.GetBySlug(ctx, solution.Slug)
		if existingSlug != nil && existingSlug.ID != solution.ID {
			return appErrors.NewConflictError("Solution with this slug already exists")
		}
		existing.Slug = sanitizeSlug(solution.Slug)
	}

	existing.Title = solution.Title
	existing.Tagline = solution.Tagline
	existing.Summary = solution.Summary
	existing.Description = solution.Description
	existing.IconURL = solution.IconURL
	existing.HeroImageURL = solution.HeroImageURL
	existing.BrochureURL = solution.BrochureURL
	existing.DisplayOrder = solution.DisplayOrder
	existing.IsFeatured = solution.IsFeatured
	existing.IsActive = solution.IsActive

	return u.solutionRepo.Update(ctx, existing)
}

func (u *solutionUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.solutionRepo.Delete(ctx, id)
}

func (u *solutionUsecase) AddFeature(ctx context.Context, feature *domain.SolutionFeature) error {
	_, err := u.solutionRepo.GetByID(ctx, feature.SolutionID)
	if err != nil {
		return err
	}
	return u.solutionRepo.AddFeature(ctx, feature)
}

func (u *solutionUsecase) UpdateFeature(ctx context.Context, feature *domain.SolutionFeature) error {
	return u.solutionRepo.UpdateFeature(ctx, feature)
}

func (u *solutionUsecase) DeleteFeature(ctx context.Context, id uuid.UUID) error {
	return u.solutionRepo.DeleteFeature(ctx, id)
}

func (u *solutionUsecase) AddModule(ctx context.Context, module *domain.SolutionModule) error {
	_, err := u.solutionRepo.GetByID(ctx, module.SolutionID)
	if err != nil {
		return err
	}
	return u.solutionRepo.AddModule(ctx, module)
}

func (u *solutionUsecase) UpdateModule(ctx context.Context, module *domain.SolutionModule) error {
	return u.solutionRepo.UpdateModule(ctx, module)
}

func (u *solutionUsecase) DeleteModule(ctx context.Context, id uuid.UUID) error {
	return u.solutionRepo.DeleteModule(ctx, id)
}
