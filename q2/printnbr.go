package PrintNbr

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var m [100]int
	i := 0
	if n < 0 {
		z01.PrintRune('-')
	}
	if n == 0 {
		z01.PrintRune('0')
	}
	n = -n
	for ; n > 0; n /= 10 {
		m[i] = n % 10
		i++
	}
	for j := i - 1; j >= 0; j-- {
		z01.PrintRune(rune(48 + m[j]))
	}
}
