package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}

	if nb == 0 || nb == 1 {
		return nb
	}

	result := 1
	for i := 1; i*i <= nb; i++ {
		result = i
	}

	if result*result == nb {
		return result
	}

	return 0
}
