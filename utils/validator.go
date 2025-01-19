package utils

import (
	"net/mail"
	"regexp"
)

// IsValidEmail uses the net/mail package to parse the email.
func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ValidatePassword(pw string) bool {
	// 1. Minimum length 8
	if len(pw) < 8 {
		return false
	}

	// 2. At least one lowercase
	hasLower, _ := regexp.MatchString("[a-z]", pw)
	if !hasLower {
		return false
	}

	// 3. At least one uppercase
	hasUpper, _ := regexp.MatchString("[A-Z]", pw)
	if !hasUpper {
		return false
	}

	// 4. At least one digit
	hasDigit, _ := regexp.MatchString("[0-9]", pw)
	if !hasDigit {
		return false
	}

	// 5. At least one special character from your chosen set
	hasSpecial, _ := regexp.MatchString(`[@$!%*?&]`, pw)
	return hasSpecial
}
