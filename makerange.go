package piscine

// import (
// 	"fmt"
// )

func AppendRange(min, max int) []int {
	var answer []int
	for i:= min ; i < max ; i++ {
		answer = append(answer, i)
	}
	return answer
}

// func main() {
// 	fmt.Println(AppendRange(5, 10))
// 	fmt.Println(AppendRange(10, 10))
// }
