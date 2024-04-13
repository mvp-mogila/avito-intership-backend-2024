package utils

import (
	"fmt"

	"github.com/mvp-mogila/avito-intership-backend-2024/internal/models"
)

type ValidationError struct {
	errorText string
}

func NewValidationError(text string) *ValidationError {
	return &ValidationError{
		errorText: text,
	}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation error: %s", e.errorText)
}

func ValidatePositive(param int, checkZero bool) error {
	if param < 0 {
		return models.ErrValidation
	}
	if checkZero && param == 0 {
		return models.ErrValidation
	}
	return nil
}
