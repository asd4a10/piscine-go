package NRune

// import (
// 	"github.com/01-edu/z01"
// )

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
// func main() {
// 	z01.PrintRune(NRune("Hello!", 3))
// 	z01.PrintRune(NRune("Salut!", 2))
// 	z01.PrintRune(NRune("Ola", 4))
// 	z01.PrintRune('\n')
// }