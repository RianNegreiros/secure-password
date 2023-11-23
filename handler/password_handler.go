// handler/password_handler.go
package handler

import (
	"encoding/json"
	"net/http"

	"github.com/RianNegreiros/secure-password/validator"
)

// PasswordRequest é a estrutura para a solicitação de senha
type PasswordRequest struct {
	Password string `json:"password"`
}

// ValidatePasswordHandler lida com as solicitações para validar senhas
func ValidatePasswordHandler(w http.ResponseWriter, r *http.Request) {
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

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
