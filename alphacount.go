package AlphaCount

// import (
// 	// "github.com/01-edu/z01"
// 	"fmt"
// )

func AlphaCount(str string) int {
	count := 0
	for index := range str {
		if str[index] >= 'A' && str[index] <= 'Z' {
			count++
		} else if str[index] >= 'a' && str[index] <= 'z' {
			count++
		}
	}
	return count
}

// func main() {
// 	str := "Hello 78 Worldadsads!    4455 /"
// 	nb := AlphaCount(str)
// 	fmt.Println(nb)
// }
