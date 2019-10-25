package piscine

// import (
// 	"fmt"
// )

func MakeRange(min, max int) []int {
	if max-min <= 0 {
		var a []int
		return a
	}
	a := make([]int, max-min)
	for i := 0; i < max-min; i++ {
		a[i] = i + min
	}
	return a
}

// func main() {
// 	fmt.Println(MakeRange(5, 10))
// 	fmt.Println(MakeRange(10, 5))
// }
