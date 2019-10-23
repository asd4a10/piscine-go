package NRune

func NRune(s string, n int) rune {
	size := 0
	for range s {
		size++
	}
	if n > size {
		return ' '
	}
	a := []rune(s)
	return a[n-1]
}
