package piscine

func SplitWhiteSpaces(s string) []string {
	work := s + " "
	var answer []string
	var word string
	for i := 0; i < len(work); i++ {
		if work[i] != ' ' {
			word = word + string(work[i])
		} else {
			if word != "" {
				answer = append(answer, word)
			}
			word = ""
		}
	}
	return answer
}
