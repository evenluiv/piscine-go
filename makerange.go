package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	answer := make([]int, max-1)
	i := min

	for i < max {
		answer[i] = i
		i++
	}
	return answer
}
