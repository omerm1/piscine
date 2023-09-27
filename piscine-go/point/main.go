package main

import (
	"github.com/01-edu/z01"
)

type Point struct {
	x, y int
}

func printRune(r rune) {
	z01.PrintRune(r)
}

func printStr(s string) {
	for _, r := range s {
		printRune(r)
	}
}

func printInt(num int) {
	// Convert the integer to a string
	numStr := ""
	if num < 0 {
		printRune('-')
		num = -num
	}

	for num > 0 {
		digit := rune('0' + num%10)
		numStr = string(digit) + numStr
		num /= 10
	}

	printStr(numStr)
}

func main() {
	points := Point{42, 21}
	printStr("x = ")
	printInt(points.x)
	printStr(", y = ")
	printInt(points.y)
	printRune('\n')
}
