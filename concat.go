package Concat

// import (
// 	"fmt"
// )

func Concat(str1 string, str2 string) string {
	var w string
	for i := range str1 {
		w += string(str1[i])
	}
	for i := range str2 {
		w += string(str2[i])
	}
	return w
}

// func main() {
// 	fmt.Println(Concat("Hello!", " How are you?"))

// }
