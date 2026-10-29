package security

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	appErrors "unotek-backend/pkg/errors"
)

type JWTClaims struct {
	UserID      uuid.UUID `json:"user_id"`
	Email       string    `json:"email"`
	Username    string    `json:"username"`
	Roles       []string  `json:"roles"`
	Permissions []string  `json:"permissions"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID uuid.UUID, email, username string, roles, permissions []string) (string, error)
	ValidateToken(tokenString string) (*JWTClaims, error)
}

type jwtService struct {
	secretKey      []byte
	expirationHours int
}

func NewJWTService(secretKey string, expirationHours int) JWTService {
	return &jwtService{
		secretKey:      []byte(secretKey),
		expirationHours: expirationHours,
	}
}

func (s *jwtService) GenerateToken(userID uuid.UUID, email, username string, roles, permissions []string) (string, error) {
	claims := &JWTClaims{
		UserID:      userID,
		Email:       email,
		Username:    username,
		Roles:       roles,
		Permissions: permissions,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(s.expirationHours) * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "unotek-backend",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(s.secretKey)
}

func (s *jwtService) ValidateToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, appErrors.ErrTokenInvalid
		}
		return s.secretKey, nil
	})

	if err != nil {
		return nil, appErrors.NewUnauthorizedError("Invalid or expired token", err)
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, appErrors.ErrTokenInvalid
}
