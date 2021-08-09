package main

import (
	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) boolean {
	if even(nbr) == 2 {
		return yes
	} else {
		return no
	}
}

func main() {
	if isEven(lengthOfArg) == 2 {
		printStr(EvenMsg)
	} else {
		printStr(OddMsg)
	}
}
