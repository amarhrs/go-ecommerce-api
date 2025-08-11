package helpers

import (
	"net/mail"
	"regexp"
	"strings"

	"amarhrs/ecommerce/models"
)

// Validate username
func ValidateUsername(username string) (bool, string) {
	username = strings.TrimSpace(username)
	if username == "" {
		return false, "Field username required"
	}
	if len(username) < 4 {
		return false, "Username must be at least 4 characters"
	}
	if len(username) > 20 {
		return false, "Username cannot be more than 20 characters"
	}
	return true, ""
}

// Validate email
func ValidateEmail(email string) (bool, string) {
	email = strings.TrimSpace(email)
	if email == "" {
		return false, "Field email required"
	}
	if _, err := mail.ParseAddress(email); err != nil {
		return false, "Invalid email format"
	}
	return true, ""
}

// ValidatePassword memvalidasi password
func ValidatePassword(password string) (bool, string) {
	password = strings.TrimSpace(password)
	if password == "" {
		return false, "Field password required"
	}
	if len(password) < 8 {
		return false, "Password must be at least 8 characters long"
	}
	if len(password) > 64 {
		return false, "Password cannot be more than 64 characters long"
	}
	if !regexp.MustCompile(`[A-Z]`).MatchString(password) {
		return false, "Password must contain at least one uppercase letter"
	}
	if !regexp.MustCompile(`[a-z]`).MatchString(password) {
		return false, "Password must contain at least one lowercase letter"
	}
	if !regexp.MustCompile(`[0-9\W]`).MatchString(password) {
		return false, "Password must contain at least one number or symbol"
	}
	return true, ""
}

// Validate login input
func ValidateLoginInput(user models.User) (bool, []string) {
	var errors []string

	if ok, msg := ValidateUsername(user.Username); !ok {
		errors = append(errors, msg)
	}

	if ok, msg := ValidatePassword(user.Password); !ok {
		errors = append(errors, msg)
	}

	return len(errors) == 0, errors
}

// Validate register input
func ValidateRegisterInput(user models.User) (bool, []string) {
	var errors []string

	if ok, msg := ValidateUsername(user.Username); !ok {
		errors = append(errors, msg)
	}

	if ok, msg := ValidateEmail(user.Email); !ok {
		errors = append(errors, msg)
	}

	if ok, msg := ValidatePassword(user.Password); !ok {
		errors = append(errors, msg)
	}

	return len(errors) == 0, errors
}
