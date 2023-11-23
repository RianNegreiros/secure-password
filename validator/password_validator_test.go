package validator

import (
	"testing"
)

func TestValidatePassword(t *testing.T) {
	tests := []struct {
		password string
		wantErr  bool
	}{
		{"Abcd123!", false},        // Valid password
		{"short", true},            // Password too short
		{"nouppercase123!", true},  // Missing uppercase letter
		{"NOLOWERCASE123!", true},  // Missing lowercase letter
		{"NoSpecialChar123", true}, // Missing special character
		{"onlylowercase", true},    // Missing uppercase, digit, and special character
		{"ONLYUPPERCASE", true},    // Missing lowercase, digit, and special character
		{"12345678", true},         // Missing uppercase, lowercase, and special character
		{"!@#$%^&*", true},         // Missing uppercase, lowercase, and digit
	}

	for _, tt := range tests {
		t.Run(tt.password, func(t *testing.T) {
			err := ValidatePassword(tt.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidatePassword(%s) error = %v, wantErr %v", tt.password, err, tt.wantErr)
				return
			}
		})
	}
}
