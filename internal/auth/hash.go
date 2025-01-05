package auth

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"time"
)

const (
	SaltLength   = 16
	PepperLength = 32
	HashLength   = 64
	Iterations   = 210000
)

const (
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 24 * time.Hour
)

func HashPassword(password string, pepper string) (string, string, error) {

	saltBytes := make([]byte, SaltLength)
	if _, err := rand.Read(saltBytes); err != nil {
		return "", "", fmt.Errorf("%w: %v", ErrSaltGeneration, err)
	}

	hashBytes := hashWithSaltAndPepper([]byte(password), saltBytes, []byte(pepper))

	hash := base64.StdEncoding.EncodeToString(hashBytes)
	salt := base64.StdEncoding.EncodeToString(saltBytes)

	return hash, salt, nil
}

func hashWithSaltAndPepper(password, salt, pepper []byte) []byte {
	pepperedPass := make([]byte, len(password)+len(pepper))
	copy(pepperedPass, password)
	copy(pepperedPass[len(password):], pepper)

	hash := hmac.New(sha512.New, salt)

	result := pepperedPass
	for i := 0; i < Iterations; i++ {
		hash.Reset()
		hash.Write(result)
		result = hash.Sum(nil)
	}

	return result
}

func VerifyPassword(password, pepper, hash, salt string) (bool, error) {

	decodedSalt, err := base64.StdEncoding.DecodeString(salt)
	if err != nil {
		return false, fmt.Errorf("error decoding salt: %w", err)
	}

	decodedHash, err := base64.StdEncoding.DecodeString(hash)
	if err != nil {
		return false, fmt.Errorf("error decoding hash: %w", err)
	}

	newHash := hashWithSaltAndPepper([]byte(password), decodedSalt, []byte(pepper))

	return subtle.ConstantTimeCompare(newHash, decodedHash) == 1, nil
}
