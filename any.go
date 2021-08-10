package piscine

func Any(f func(string) bool, a []string) bool {
	for _, data := range a {
		if f(data) == true {
			return true
		} else {
			continue
		}
	}
	return false
}
