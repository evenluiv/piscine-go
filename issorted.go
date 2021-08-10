package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	for i, j := 0, 1; i < len(a)-1; i, j = i+1, j+1 {
		if a[i] < a[j] {
			continue
		} else {
			return false
		}
	}
	return true
}
