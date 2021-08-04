package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, rune1 := range s {
		z01.PrintRune(rune1)
	}
}
