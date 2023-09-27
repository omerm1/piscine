package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	wordStart := -1

	for i, char := range s {
		if char == ' ' || char == '\t' || char == '\n' {
			if wordStart != -1 {
				result = append(result, s[wordStart:i])
				wordStart = -1
			}
		} else if wordStart == -1 {
			wordStart = i
		}
	}

	if wordStart != -1 {
		result = append(result, s[wordStart:])
	}

	return result
}
