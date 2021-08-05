package piscine

func MakeRange(min, max int) []int {
	if max <= min {
		return nil
	}

	answer := make([]int, max-min)
	i := 0

	for i < max-min {
		answer[i] = i + min
		i++
	}
	return answer
}
