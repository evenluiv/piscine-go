package piscine

func ForEach(f func(int), a []int) {
	for _, data := range a {
		f(data)
	}
}
