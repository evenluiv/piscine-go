package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))

	for index, data := range a {
		result[index] = f(data)
	}

	return result
}
