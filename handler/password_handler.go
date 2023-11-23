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

	// Check Content-Type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "O tipo de conteúdo deve ser application/json", http.StatusUnsupportedMediaType)
		return
	}

	var req PasswordRequest
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)
	if err != nil {
		http.Error(w, "Erro ao decodificar a solicitação JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = validator.ValidatePassword(req.Password)
	if err != nil {
		http.Error(w, "Erro na validação da senha: "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
