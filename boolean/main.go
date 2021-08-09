package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func lengthOfArg() int {
	args := os.Args
	lenOfArgs := len(args) - 1
	return lenOfArgs
}

func EvenMsg() {
	printStr("I have an even number of arguments")
}

func OddMsg() {
	printStr("I have an odd number of arguments")
}

func main() {
	if isEven(lengthOfArg()) == true {
		EvenMsg()
	} else {
		OddMsg()
	}
}
