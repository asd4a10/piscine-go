package BasicAtoi2

// import (
// 	"fmt"
// )

func BasicAtoi2(s string) int {
	size := 0
	for index := range s {
		if int(s[index]-'0') < 0 || int(s[index]-'0') > 9 {
			return 0
		}
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
// 	s3 := "012 345"
// 	s4 := "Hello World!"

// 	n := BasicAtoi2(s)
// 	n2 := BasicAtoi2(s2)
// 	n3 := BasicAtoi2(s3)
// 	n4 := BasicAtoi2(s4)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// 	fmt.Println(n4)

// }
