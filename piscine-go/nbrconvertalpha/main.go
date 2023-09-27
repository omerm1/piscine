package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	upper := false
	args := os.Args[1:]

	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	for _, arg := range args {
		pos := atoi(arg)
		if pos < 1 || pos > 26 {
			z01.PrintRune(' ')
		} else {
			letter := intToLatinLetter(pos, upper)
			z01.PrintRune(letter)
		}
	}

	z01.PrintRune('\n')
}

func atoi(s string) int {
	num := 0

	for _, c := range s {
		if c < '0' || c > '9' {
			return 0
		}
		num = num*10 + int(c-'0')
	}

	return num
}

func intToLatinLetter(pos int, upper bool) rune {
	offset := pos - 1

	if upper {
		return rune('A' + offset)
	}

	return rune('a' + offset)
}
