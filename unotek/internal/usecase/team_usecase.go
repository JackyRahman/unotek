package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	"unotek-backend/pkg/pagination"
)

type teamUsecase struct {
	teamRepo domain.TeamRepository
}

func NewTeamUsecase(teamRepo domain.TeamRepository) domain.TeamUsecase {
	return &teamUsecase{teamRepo: teamRepo}
}

func (u *teamUsecase) List(ctx context.Context, params pagination.Params, publicOnly bool) ([]domain.TeamMember, int64, error) {
	return u.teamRepo.List(ctx, params, publicOnly)
}

func (u *teamUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.TeamMember, error) {
	return u.teamRepo.GetByID(ctx, id)
}

func (u *teamUsecase) Create(ctx context.Context, member *domain.TeamMember) error {
	return u.teamRepo.Create(ctx, member)
}

func (u *teamUsecase) Update(ctx context.Context, member *domain.TeamMember) error {
	existing, err := u.teamRepo.GetByID(ctx, member.ID)
	if err != nil {
		return err
	}

	existing.FullName = member.FullName
	existing.Position = member.Position
	existing.Department = member.Department
	existing.Bio = member.Bio
	existing.PhotoURL = member.PhotoURL
	existing.LinkedinURL = member.LinkedinURL
	existing.Email = member.Email
	existing.DisplayOrder = member.DisplayOrder
	existing.IsLeadership = member.IsLeadership
	existing.IsActive = member.IsActive

	return u.teamRepo.Update(ctx, existing)
}

func (u *teamUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.teamRepo.Delete(ctx, id)
}
