package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	"unotek-backend/pkg/pagination"
)

type contactUsecase struct {
	contactRepo domain.ContactRepository
}

func NewContactUsecase(contactRepo domain.ContactRepository) domain.ContactUsecase {
	return &contactUsecase{contactRepo: contactRepo}
}

func (u *contactUsecase) Submit(ctx context.Context, contact *domain.ContactSubmission) error {
	contact.Status = "new"
	return u.contactRepo.Create(ctx, contact)
}

func (u *contactUsecase) List(ctx context.Context, params pagination.Params, statusFilter string) ([]domain.ContactSubmission, int64, error) {
	return u.contactRepo.List(ctx, params, statusFilter)
}

func (u *contactUsecase) GetByID(ctx context.Context, id uuid.UUID) (*domain.ContactSubmission, error) {
	return u.contactRepo.GetByID(ctx, id)
}

func (u *contactUsecase) UpdateStatus(ctx context.Context, id uuid.UUID, status, notes string) error {
	return u.contactRepo.UpdateStatus(ctx, id, status, notes)
}

func (u *contactUsecase) Delete(ctx context.Context, id uuid.UUID) error {
	return u.contactRepo.Delete(ctx, id)
}
