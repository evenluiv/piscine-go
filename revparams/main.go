package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for i := len(arg) - 1; i > 0; i-- {
		for _, a := range arg[i] {
			if i < len(arg) {
				z01.PrintRune(a)
			}
		}
		if i > 0 {
			z01.PrintRune('\n')
		}
	}
}
