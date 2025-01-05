package models

import "time"

type AuthUser struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,passwordComplexity"`
}

type AuthTokenPair struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type AuthRequest struct {
	Data *AuthUser `json:"data"`
}

type AuthResponse struct {
	Data *AuthTokenPair `json:"data"`
}

type User struct {
	ID           int64
	Email        string
	PasswordHash string
	Salt         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type AccessTokenRefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

type AccessTokenRequest struct {
	Data *AccessTokenRefreshTokenRequest `json:"data"`
}

type AccessTokenResponse struct {
	Data *AuthTokenPair `json:"data"`
}
