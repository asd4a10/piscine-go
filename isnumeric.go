package IsNumeric

// import (
// 	"fmt"
// )

func IsNumeric(a string) bool {
	for i := range a {
		if a[i] > '9' || a[i] < '0' {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(IsAlpha("Hello! How are you?"))
// 	fmt.Println(IsAlpha("HelloHowareyou"))
// 	fmt.Println(IsAlpha("What's this 4?"))
// 	fmt.Println(IsAlpha("Whatsthis4"))

// }
