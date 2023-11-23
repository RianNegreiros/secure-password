package main

import (
	"encoding/json"
	"net/http"
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

	w.WriteHeader(http.StatusNoContent)
}
