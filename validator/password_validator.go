// validator/password_validator.go
package validator

import (
	"fmt"
	"regexp"
	"strings"
)

// ValidatePassword verifica se a senha atende aos critérios de segurança
func ValidatePassword(password string) error {
	var errors []string

	// Verificar o comprimento mínimo da senha
	if len(password) < 8 {
		errors = append(errors, "A senha deve ter pelo menos 8 caracteres")
	}

	// Verificar se contém pelo menos uma letra maiúscula
	hasUpperCase, _ := regexp.MatchString(`[A-Z]`, password)
	if !hasUpperCase {
		errors = append(errors, "A senha deve conter pelo menos uma letra maiúscula")
	}

	// Verificar se contém pelo menos uma letra minúscula
	hasLowerCase, _ := regexp.MatchString(`[a-z]`, password)
	if !hasLowerCase {
		errors = append(errors, "A senha deve conter pelo menos uma letra minúscula")
	}

	// Verificar se contém pelo menos um dígito numérico
	hasDigit, _ := regexp.MatchString(`\d`, password)
	if !hasDigit {
		errors = append(errors, "A senha deve conter pelo menos um dígito numérico")
	}

	// Verificar se contém pelo menos um caractere especial
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%]`, password)
	if !hasSpecialChar {
		errors = append(errors, "A senha deve conter pelo menos um caractere especial (!@#$%)")
	}

	if len(errors) > 0 {
		return fmt.Errorf(strings.Join(errors, "\n"))
	}

	return nil
}
