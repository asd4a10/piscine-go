package IsAlpha

// import (
// 	"fmt"
// )

func IsAlpha(a string) bool {
	for i := range a {
		if (a[i] < 'a' || a[i] > 'z') && (a[i] < 'A' || a[i] > 'Z') && (a[i] > '9' || a[i] < '0') {
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
