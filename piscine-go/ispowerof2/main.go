package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPowerOfTwo(n int) string {
	for n > 1 {
		if n%2 != 0 {
			return "false"
		}
		n /= 2
	}
	return "true"
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}

		result := isPowerOfTwo(num)
		for _, char := range result {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
