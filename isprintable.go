package piscine

func IsPrintable(s string) bool {
	string := []byte(s)
	for _, value := range string {
		if value > 31 {
			continue
		} else {
			return false
		}
	}
	return true
}
