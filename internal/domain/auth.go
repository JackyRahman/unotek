package domain

import (
	"context"

	"github.com/google/uuid"
)

type AuthUser struct {
	ID          uuid.UUID `json:"id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	FullName    string    `json:"full_name"`
	AvatarURL   string    `json:"avatar_url"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
}

type AuthResponse struct {
	AccessToken string   `json:"access_token"`
	TokenType   string   `json:"token_type"`
	ExpiresIn   int      `json:"expires_in"` // in seconds
	User        AuthUser `json:"user"`
}

type AuthUsecase interface {
	Login(ctx context.Context, identifier, password string) (*AuthResponse, error)
	GetProfile(ctx context.Context, userID uuid.UUID) (*AuthUser, error)
}
