package ToLower

// import (
// 	"fmt"
// )

func ToLower(s string) string {
	a := []rune(s)
	for i := range a {
		if a[i] >= 'A' && a[i] <= 'Z' {
			a[i] += 32
		}
	}
	s = string(a)
	return s
}

// func main() {
// 	fmt.Println(ToUpper("Hello! How are you?"))
// }
