package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args[1:]

	sortArguments(arguments)

	for _, arg := range arguments {
		for _, char := range arg {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}

func sortArguments(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if compare(args[i], args[j]) > 0 {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func compare(str1, str2 string) int {
	for i := 0; i < len(str1) && i < len(str2); i++ {
		if str1[i] < str2[i] {
			return -1
		} else if str1[i] > str2[i] {
			return 1
		}
	}

	if len(str1) < len(str2) {
		return -1
	} else if len(str1) > len(str2) {
		return 1
	}
	return 0
}
