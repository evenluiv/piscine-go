package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for index, s := range arg[0] {
		if index > 1 {
			z01.PrintRune(s)
		}
	}
	z01.PrintRune('\n')
}
