package auth

import "errors"

type AuthTokenPair struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserData struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

var (
	ErrInvalidToken       = errors.New("invalid token")
	ErrTokenExpired       = errors.New("token has expired")
	ErrInvalidSignature   = errors.New("invalid token signature")
	ErrSaltGeneration     = errors.New("error generating salt")
	ErrInvalidCredentials = errors.New("invalid credentials")
)
