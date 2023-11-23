package main

import (
	"net/http"

	"github.com/RianNegreiros/secure-password/handler"
)

func main() {
	http.HandleFunc("/validate-password", handler.ValidatePasswordHandler)
	http.ListenAndServe(":8080", nil)
}
