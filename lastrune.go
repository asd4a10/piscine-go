package NRune

func NRune(s string) rune {
	size := 0
	for range s {
		size++
	}
	// if n > size || n < 1 {
	// 	return '\x00'
	// }
	a := []rune(s)
	return a[size-1]
}
