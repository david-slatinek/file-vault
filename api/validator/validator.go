package validator

import "unicode"

var InvalidPassword = "invalid password. Must contain at least 1 uppercase letter, 1 lowercase letter, 1 digit, and 1 special character"

func ValidatePassword(password string) bool {
	var uppercase, lowercase, digit, specialCharacter = false, false, false, false

	for _, r := range password {
		if unicode.IsUpper(r) && unicode.IsLetter(r) {
			uppercase = true
		}

		if unicode.IsLower(r) && unicode.IsLetter(r) {
			lowercase = true
		}

		if unicode.IsDigit(r) {
			digit = true
		}

		if unicode.IsPunct(r) || unicode.IsSymbol(r) {
			specialCharacter = true
		}

		if uppercase && lowercase && digit && specialCharacter {
			return true
		}
	}

	return false
}
