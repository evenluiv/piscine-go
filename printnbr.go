package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(a int) {
	n := []int{}
	if a < 0 {
		z01.PrintRune('-')
		a = -a
	}
	if a == 0 {
		z01.PrintRune('0')
	}

	for a != 0 {
		n = append(n, a%10)
		b := a % 10
		if a > 0 {
			a = (a - b) / 10
		} else if a < 0 {
			a = (a + b) / 10
		}
	}

	p := len(n)
	for p > 0 {
		if n[p-1] < 0 {
			n[p-1] = -n[p-1]
		}
		z01.PrintRune(rune(n[p-1] + 48))
		p--
	}
}
