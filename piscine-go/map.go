package piscine

func Map(f func(int) bool, a []int) []bool {
	tab := make([]bool, len(a))

	for k, v := range a {
		tab[k] = f(v)
	}
	return tab
}
