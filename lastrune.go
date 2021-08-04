package piscine

func LastRune(s string) rune {
	string := []rune(s)

	return string[len(string)-1]
}
