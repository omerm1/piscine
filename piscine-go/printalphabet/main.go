package main

import "github.com/01-edu/z01"

func main() {
	for letter := 'a'; letter <= 'z'; letter++ {
		z01.PrintRune(letter)
	}
	z01.PrintRune('\n')
}
