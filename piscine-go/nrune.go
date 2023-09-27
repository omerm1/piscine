package piscine

func NRune(s string, n int) rune {
	if n <= 0 || n > len(s) {
		return 0
	}

	count := 0
	for _, r := range s {
		if count == n-1 {
			return r
		}
		count++
	}

	return 0
}
