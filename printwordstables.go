package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, b := range a {
		SplitWhiteSpaces(b)
		z01.PrintRune('\n')
	}
}
