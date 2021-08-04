package piscine

func ToUpper(s string) string {
	arr := []byte(s)
	for index, value := range arr {
		if value >= 'a' && value <= 'z' {
			arr[index] = value - 32
		}
	}
	return string(arr)
}
