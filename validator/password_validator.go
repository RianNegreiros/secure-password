// validator/password_validator.go
package validator

import (
	"fmt"
	"regexp"
)

// ValidatePassword verifica se a senha atende aos critérios de segurança
func ValidatePassword(password string) error {
	// Verificar o comprimento mínimo da senha
	if len(password) < 8 {
		return fmt.Errorf("A senha deve ter pelo menos 8 caracteres")
	}

	// Verificar se contém pelo menos uma letra maiúscula
	hasUpperCase, _ := regexp.MatchString(`[A-Z]`, password)
	if !hasUpperCase {
		return fmt.Errorf("A senha deve conter pelo menos uma letra maiúscula")
	}

	// Verificar se contém pelo menos uma letra minúscula
	hasLowerCase, _ := regexp.MatchString(`[a-z]`, password)
	if !hasLowerCase {
		return fmt.Errorf("A senha deve conter pelo menos uma letra minúscula")
	}

	// Verificar se contém pelo menos um dígito numérico
	hasDigit, _ := regexp.MatchString(`\d`, password)
	if !hasDigit {
		return fmt.Errorf("A senha deve conter pelo menos um dígito numérico")
	}

	// Verificar se contém pelo menos um caractere especial
	hasSpecialChar, _ := regexp.MatchString(`[!@#$%]`, password)
	if !hasSpecialChar {
		return fmt.Errorf("A senha deve conter pelo menos um caractere especial (!@#$%)")
	}

	return nil
}
