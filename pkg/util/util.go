package util

import (
	"regexp"
	"unicode"
)

var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func IsValidEmail(email string) bool {
	return emailRegex.MatchString(email)
}

func ValidatePassword(password string) (string, bool) {
	var (
		hasUppercase = false
		hasLowercase = false
		hasNumber    = false
		hasSpecial   = false
		// specialRunes = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
	)

	if len(password) < 8 {
		return "Password must be at least 8 characters long.", false
	}

	for _, char := range password {
		switch {
			case unicode.IsUpper(char):
				hasUppercase = true
			case unicode.IsLower(char):
				hasLowercase = true
			case unicode.IsNumber(char):
				hasNumber = true
			case unicode.IsPunct(char) || unicode.IsSymbol(char):
				hasSpecial = true

			
		}
	}

	if !hasUppercase {
		return "Password must contain at least one uppercase letter.", false
	}

	if !hasLowercase {
		return "Password must contain at least one lowercase letter.", false
	}

	if !hasNumber {
		return "Password must contain at least one number.", false
	}

	if !hasSpecial {
		return "Password must contain at least one special character.", false
	}

	return "", true
}