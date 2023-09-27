package piscine

func Index(s, toFind string) int {
	sLen := len(s)
	toFindLen := len(toFind)

	if toFindLen == 0 {
		return 0
	}

	for i := 0; i <= sLen-toFindLen; i++ {
		if s[i:i+toFindLen] == toFind {
			return i
		}
	}

	return -1
}
