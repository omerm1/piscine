package piscine

func Capitalize(s string) string {
	runes := []rune(s)
	isFirstLetter := true
	for i := range runes {
		if runes[i] >= 'a' && runes[i] <= 'z' && isFirstLetter {
			runes[i] -= 32
			isFirstLetter = false
		} else if (IsRuneDigit(runes[i]) || (runes[i] >= 'A' && runes[i] <= 'Z')) && isFirstLetter {
			isFirstLetter = false
		} else if !isFirstLetter && runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] += 32
		} else if !isFirstLetter && !IsRuneDigit(runes[i]) && !IsRuneAlpha(runes[i]) {
			isFirstLetter = true
		}
	}

	return string(runes)
}

func IsRuneAlpha(r rune) bool {
	if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
		return true
	}

	return false
}

func IsRuneDigit(r rune) bool {
	if r >= '0' && r <= '9' {
		return true
	}

	return false
}
