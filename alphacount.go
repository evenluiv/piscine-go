package piscine

func AlphaCount(s string) int {
	string := []byte(s)
	counter := 0
	for _, bytevalue := range string {
		if bytevalue >= 65 && bytevalue <= 90 || bytevalue >= 97 && bytevalue <= 122 {
			counter++
		} else {
			continue
		}
	}
	return counter
}
