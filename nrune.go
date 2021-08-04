package piscine

func NRune(s string, n int) rune {
	a := n - 1
	if a >= 0 && a < len(s) {
		for index, letter := range s {
			if index == a {
				return letter
			}
		}
	}
	return 0
}
