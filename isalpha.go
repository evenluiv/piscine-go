package piscine

func IsAlpha(s string) bool {
	string := []byte(s)
	for _, value := range string {
		if value >= 'a' && value <= 'z' || value >= 'A' && value <= 'Z' || value >= '0' && value <= '9' {
			continue
		} else {
			return false
		}
	}
	return true
}
