package utils

func HasDigit(input string) bool {
	for _, c := range input {
		if c >= '0' && c <= 9 {
			return true
		}
	}
	return false
}

func HasLetter(input string) bool {
	for _, c := range input {
		if c < '0' || c > '9' {
			return true
		}
	}
	return false
}
