package piscine

func IsUpper(s string) bool {
	string := []byte(s)
	for _, value := range string {
		if value >= 'A' && value <= 'Z' {
			continue
		} else {
			return false
		}
	}
	return true
}
