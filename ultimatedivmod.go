package piscine

func UltimateDivMod(a *int, b *int) {
	g := *a
	h := *b
	*a = g / h
	*b = g % h
}
