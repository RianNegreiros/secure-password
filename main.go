package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

// PasswordRequest é a estrutura para a solicitação de senha
type PasswordRequest struct {
	Password string `json:"password"`
}

func main() {
	http.HandleFunc("/validate-password", validatePasswordHandler)
	http.ListenAndServe(":8080", nil)
}

func validatePasswordHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var req PasswordRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar a solicitação JSON", http.StatusBadRequest)
		return
	}

	err = validatePassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func validatePassword(password string) error {
	// Verificar o comprimento mínimo da senha
	if len(password) < 8 {
		return fmt.Errorf("A senha deve ter pelo menos 8 caracteres")
	}

	// Verificar se contém pelo menos uma letra maiúscula
	if strings.ToUpper(password) == password {
		return fmt.Errorf("A senha deve conter pelo menos uma letra maiúscula")
	}

	// Verificar se contém pelo menos uma letra minúscula
	if strings.ToLower(password) == password {
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
