package piscine

func IsNumeric(s string) bool {
	string := []byte(s)
	for _, value := range string {
		if value >= 33 && value <= 126 {
			continue
		} else {
			return false
		}
	}
	return true
}
