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

// ValidatePassword verifica se a senha atende aos critérios de segurança
func ValidatePassword(password string) error {
	var errors []string

	// Verificar o comprimento mínimo da senha
	if len(password) < minLength {
		errors = append(errors, errMinLength)
	}

	// Verificar se contém pelo menos uma letra maiúscula
	hasUpperCase, _ := regexp.MatchString(upperCaseRegex, password)
	if !hasUpperCase {
		errors = append(errors, errUpperCase)
	}

	// Verificar se contém pelo menos uma letra minúscula
	hasLowerCase, _ := regexp.MatchString(lowerCaseRegex, password)
	if !hasLowerCase {
		errors = append(errors, errLowerCase)
	}

	// Verificar se contém pelo menos um dígito numérico
	hasDigit, _ := regexp.MatchString(digitRegex, password)
	if !hasDigit {
		errors = append(errors, errDigit)
	}

	// Verificar se contém pelo menos um caractere especial
	hasSpecialChar, _ := regexp.MatchString(specialCharRegex, password)
	if !hasSpecialChar {
		errors = append(errors, errSpecialChar)
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
}
