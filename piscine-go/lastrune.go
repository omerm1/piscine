package piscine

func LastRune(s string) rune {
	var lastRune rune

	for _, r := range s {
		lastRune = r
	}

	return lastRune
}
