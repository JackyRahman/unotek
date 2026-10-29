package usecase

import (
	"context"

	"github.com/google/uuid"
	"unotek-backend/internal/domain"
	appErrors "unotek-backend/pkg/errors"
	"unotek-backend/pkg/security"
)

type authUsecase struct {
	userRepo   domain.UserRepository
	rbacRepo   domain.RBACRepository
	jwtService security.JWTService
}

func NewAuthUsecase(userRepo domain.UserRepository, rbacRepo domain.RBACRepository, jwtService security.JWTService) domain.AuthUsecase {
	return &authUsecase{
		userRepo:   userRepo,
		rbacRepo:   rbacRepo,
		jwtService: jwtService,
	}
}

func (u *authUsecase) Login(ctx context.Context, identifier, password string) (*domain.AuthResponse, error) {
	user, err := u.userRepo.GetByUsernameOrEmail(ctx, identifier)
	if err != nil {
		return nil, appErrors.ErrInvalidCredentials
	}

	if !user.IsActive {
		return nil, appErrors.NewForbiddenError("User account is inactive. Please contact administrator.")
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		return nil, appErrors.ErrInvalidCredentials
	}

	var roles []string
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}

	permissions, err := u.rbacRepo.GetPermissionsByRoleNames(ctx, roles)
	if err != nil {
		return nil, err
	}

	token, err := u.jwtService.GenerateToken(user.ID, user.Email, user.Username, roles, permissions)
	if err != nil {
		return nil, err
	}

	// Update last login
	_ = u.userRepo.UpdateLastLogin(ctx, user.ID)

	return &domain.AuthResponse{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   86400, // 24 hours
		User: domain.AuthUser{
			ID:          user.ID,
			Email:       user.Email,
			Username:    user.Username,
			FullName:    user.FullName,
			AvatarURL:   user.AvatarURL,
			Roles:       roles,
			Permissions: permissions,
		},
	}, nil
}

func (u *authUsecase) GetProfile(ctx context.Context, userID uuid.UUID) (*domain.AuthUser, error) {
	user, err := u.userRepo.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}

	var roles []string
	for _, r := range user.Roles {
		roles = append(roles, r.Name)
	}

	permissions, err := u.rbacRepo.GetPermissionsByRoleNames(ctx, roles)
	if err != nil {
		return nil, err
	}

	return &domain.AuthUser{
		ID:          user.ID,
		Email:       user.Email,
		Username:    user.Username,
		FullName:    user.FullName,
		AvatarURL:   user.AvatarURL,
		Roles:       roles,
		Permissions: permissions,
	}, nil
}
