package piscine

func CountIf(f func(string) bool, tab []string) int {
	counter := 0

	for _, data := range tab {
		if f(data) == true {
			counter++
		} else {
			continue
		}
	}
	return counter
}
