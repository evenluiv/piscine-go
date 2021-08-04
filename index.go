package piscine

func Index(s string, toFind string) int {
	string := []byte(s)
	n := len(toFind)
	for index := range string {
		if index+n > len(string) {
			return -1
		} else {
			if s[index:index+n] == toFind {
				return index
			}
		}
	}
	return -1
}
