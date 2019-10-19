package Atoi

// import (
// 	"fmt"
// )

func Atoi(s string) int {
	size := 0
	ind := '0'
	for index := range s {
		size++
		if index == 0 && s[index] == '+' {
			ind = '+'
		} else if index == 0 && s[index] == '-' {
			ind = '-'
		} else if int(s[index]-'0') < 0 || int(s[index]-'0') > 9 {
			return 0
		}
	}
	d := 1
	num := 0
	if ind == '+' {
		for i := size - 1; i >= 1; i-- {
			num += d * int(s[i]-'0')
			d *= 10
		}
		return num
	} else if ind == '-' {
		for i := size - 1; i >= 1; i-- {
			num += d * int(s[i]-'0')
			d *= 10
		}
		return -num
	} else {
		for i := size - 1; i >= 0; i-- {
			num += d * int(s[i]-'0')
			d *= 10
		}
		return num
	}
}

// func main() {
// 	// s := "12345"
// 	s2 := "0000000012345"
// 	s3 := "012345"
// 	s4 := "Hello World!"
// 	s5 := "+1234"
// 	s6 := "-1234"
// 	s7 := "++1234"
// 	s8 := "--1234"

// 	// n := Atoi(s)
// 	n2 := Atoi(s2)
// 	n3 := Atoi(s3)
// 	n4 := Atoi(s4)
// 	n5 := Atoi(s5)
// 	n6 := Atoi(s6)
// 	n7 := Atoi(s7)
// 	n8 := Atoi(s8)

// 	fmt.Println('+' - '0')
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// 	fmt.Println(n4)
// 	fmt.Println(n5)
// 	fmt.Println(n6)
// 	fmt.Println(n7)
// 	fmt.Println(n8)
// }
