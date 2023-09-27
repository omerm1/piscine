package piscine

func IsAlpha(s string) bool {
	for _, char := range s {
		if !('a' <= char && char <= 'z') && !('A' <= char && char <= 'Z') && !('0' <= char && char <= '9') {
			return false
		}
	}

	return true
}
