package piscine

func Rot14(s string) string {
	result := []rune(s)

	for i, char := range result {
		if ('a' <= char && char <= 'z') || ('A' <= char && char <= 'Z') {
			base := 'a'
			if 'A' <= char && char <= 'Z' {
				base = 'A'
			}

			result[i] = base + (char-base+14)%26
		}
	}
	return string(result)
}
