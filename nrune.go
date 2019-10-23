package NRune

// import (
// 	"github.com/01-edu/z01"
// )

func NRune(s string, n int) rune {
	a := []rune(s)
	return a[n-1]
}

// func main() {
// 	z01.PrintRune(FirstRune("Hello!"))
// 	z01.PrintRune(FirstRune("Salut!"))
// 	z01.PrintRune(FirstRune("Ola!"))
// 	z01.PrintRune('\n')
// }
