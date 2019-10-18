package BasicAtoi

// import (
// 	"fmt"
// )

func BasicAtoi(s string) int {
	size := 0
	for range s {
		size++
	}
	d := 1
	num := 0
	for i := size - 1; i >= 0; i-- {
		num += d * int(s[i]-'0')
		d *= 10
	}

	return num
}

// func main() {
// 	s := "12345"
// 	s2 := "0000000012345"
// 	s3 := "000000"

// 	n := BasicAtoi(s)
// 	n2 := BasicAtoi(s2)
// 	n3 := BasicAtoi(s3)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// }
