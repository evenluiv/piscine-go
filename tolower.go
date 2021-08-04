package piscine

func ToLower(s string) string {
	arr := []byte(s)
	for index, value := range arr {
		if value >= 'A' && value <= 'Z' {
			arr[index] = value + 32
		}
	}
	return string(arr)
}
