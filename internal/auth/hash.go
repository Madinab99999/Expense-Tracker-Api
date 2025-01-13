package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	SaltLength           = 16
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 24 * time.Hour
	BcryptCost           = bcrypt.DefaultCost
)

func HashPassword(password string) (string, string, error) {

	saltBytes := make([]byte, SaltLength)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrSaltGeneration, err)
	}

	saltedPassword := append([]byte(password), saltBytes...)
	hashedPassword, err := bcrypt.GenerateFromPassword(saltedPassword, BcryptCost)
	if err != nil {
		return "", "", fmt.Errorf("failed to hash password: %w", err)
	}

	hash := base64.StdEncoding.EncodeToString(hashedPassword)
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	return hash, salt, nil
}

func VerifyPassword(password, hash, salt string) (bool, error) {

	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false, fmt.Errorf("error decoding salt: %w", err)
	}

	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false, fmt.Errorf("error decoding hash: %w", err)
	}

	saltedPassword := append([]byte(password), decodedSalt...)

	if err := bcrypt.CompareHashAndPassword(decodedHash, saltedPassword); err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		}
		return false, fmt.Errorf("error comparing hash and password: %w", err)
	}

	return true, nil
}
