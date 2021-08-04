package piscine

func IterativeFactorial(nb int) int {
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 20 {
		return 0
	}

	j := nb - 1

	for i := j; i >= 1; i-- {
		nb = nb * i
	}

	return nb
}
