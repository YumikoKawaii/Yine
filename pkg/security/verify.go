package security

import "unicode"

func VerifyPassword(p string) bool {

	length := len(p) >= 8
	number := false
	upper := false
	special := false

	for _, c := range p {
		switch {
		case unicode.IsNumber(c):
			number = true
		case unicode.IsUpper(c):
			upper = true
		case unicode.IsPunct(c) || unicode.IsSymbol(c):
			special = true
		}

	}

	return length && number && upper && special
}
