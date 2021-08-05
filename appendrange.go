package piscine

func AppendRange(min, max int) []int {
	if max <= min {
		return nil
	}

	var answer []int
	i := min

	for i < max {
		answer = append(answer, i)
		i++
	}
	return answer
}
