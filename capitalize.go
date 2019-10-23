package Capitalize

// import (
// 	"fmt"
// )

func Capitalize(s string) string {
	a := []rune(s)
	w := 0
	for i := range a {
		if (a[i] >= 'A' && a[i] <= 'Z') || (a[i] >= 'a' && a[i] <= 'z') {
			w++
		} else {
			w = 0
		}

		if w == 1 && a[i] >= 'a' && a[i] <= 'z' {
			a[i] -= 32
		} else if w > 1 && a[i] >= 'A' && a[i] <= 'Z' {
			a[i] += 32
		}
	}
	s = string(a)
	return s
}

// func main() {
// 	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
// }
