package piscine

func IsLower(s string) bool {
	string := []byte(s)
	for _, value := range string {
		if value >= 'a' && value <= 'z' {
			continue
		} else {
			return false
		}
	}
	return true
}
