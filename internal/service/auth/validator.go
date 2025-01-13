package auth_service

import (
	"fmt"
	"regexp"

	"github.com/Madinab99999/Expense-Tracker-Api/internal/models"

	"github.com/go-playground/validator/v10"
)

func validatePasswordComplexity(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasNumber && hasSpecial && len(password) >= 8
}

func (s *AuthService) validateAuthUser(user *models.AuthUser) error {
	if err := s.validator.Struct(user); err != nil {
		return fmt.Errorf("%w: %v", ErrValidation, err)
	}
	return nil
}
