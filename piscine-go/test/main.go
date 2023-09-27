package main

import (
	"fmt"
	"piscine"
)

func main() {
	steps := piscine.CollatzCountdown(20)
	fmt.Println(steps)
}
