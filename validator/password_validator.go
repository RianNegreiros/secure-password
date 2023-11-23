package validator

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	minLength        = 8
	upperCaseRegex   = `[A-Z]`
	lowerCaseRegex   = `[a-z]`
	digitRegex       = `\d`
	specialCharRegex = `[!@#$%]`
)

const (
	errMinLength   = "A senha deve ter pelo menos 8 caracteres"
	errUpperCase   = "A senha deve conter pelo menos uma letra maiúscula"
	errLowerCase   = "A senha deve conter pelo menos uma letra minúscula"
	errDigit       = "A senha deve conter pelo menos um dígito numérico"
	errSpecialChar = "A senha deve conter pelo menos um caractere especial (!@#$%)"
)

// ValidatePassword checks that the password meets the security criteria
func ValidatePassword(password string) error {
	var errors []string

	// Check the minimum password length
	if len(password) < minLength {
		errors = append(errors, errMinLength)
	}

	// Check that it contains at least one capital letter
	hasUpperCase, _ := regexp.MatchString(upperCaseRegex, password)
	if !hasUpperCase {
		errors = append(errors, errUpperCase)
	}

	// Check that it contains at least one lowercase letter
	hasLowerCase, _ := regexp.MatchString(lowerCaseRegex, password)
	if !hasLowerCase {
		errors = append(errors, errLowerCase)
	}

	// Check that it contains at least one numeric digit
	hasDigit, _ := regexp.MatchString(digitRegex, password)
	if !hasDigit {
		errors = append(errors, errDigit)
	}

	// Check if it contains at least one special character
	hasSpecialChar, _ := regexp.MatchString(specialCharRegex, password)
	if !hasSpecialChar {
		errors = append(errors, errSpecialChar)
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
}
