package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestValidatePasswordHandler(t *testing.T) {
	// Test valid JSON request with a valid password
	validRequest := `{"password": "SecurePwd123!"}`
	validRequestReader := bytes.NewBufferString(validRequest)
	validRequestRec := httptest.NewRequest(http.MethodPost, "/validate", validRequestReader)
	validRequestRec.Header.Set("Content-Type", "application/json")
	validRequestResp := httptest.NewRecorder()

	ValidatePasswordHandler(validRequestResp, validRequestRec)

	if validRequestResp.Code != http.StatusNoContent {
		t.Errorf("Expected status code %d, got %d", http.StatusNoContent, validRequestResp.Code)
	}

	// Test invalid HTTP method
	invalidMethodRec := httptest.NewRequest(http.MethodGet, "/validate", nil)
	invalidMethodResp := httptest.NewRecorder()

	ValidatePasswordHandler(invalidMethodResp, invalidMethodRec)

	if invalidMethodResp.Code != http.StatusMethodNotAllowed {
		t.Errorf("Expected status code %d, got %d", http.StatusMethodNotAllowed, invalidMethodResp.Code)
	}

	// Test invalid Content-Type
	invalidContentTypeRequest := `{"password": "WeakPwd"}`
	invalidContentTypeRequestReader := bytes.NewBufferString(invalidContentTypeRequest)
	invalidContentTypeRequestRec := httptest.NewRequest(http.MethodPost, "/validate", invalidContentTypeRequestReader)
	invalidContentTypeRequestRec.Header.Set("Content-Type", "text/plain")
	invalidContentTypeResp := httptest.NewRecorder()

	ValidatePasswordHandler(invalidContentTypeResp, invalidContentTypeRequestRec)

	if invalidContentTypeResp.Code != http.StatusUnsupportedMediaType {
		t.Errorf("Expected status code %d, got %d", http.StatusUnsupportedMediaType, invalidContentTypeResp.Code)
	}
}
