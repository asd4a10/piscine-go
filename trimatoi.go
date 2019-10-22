package TrimAtoi

// import (
// 	"fmt"
// )

func TrimAtoi(w string) int {
	size := 0
	ind := '0'
	var s string
	for p := range w {
		if rune(w[p]) >= '0' && rune(w[p]) <= '9' {
			break
		}
		if rune(w[p]) == '+' {
			ind = '+'
			break
		}
		if rune(w[p]) == '-' {
			ind = '-'
			break
		}
	}
	for p := range w {
		if rune(w[p]) >= '0' && rune(w[p]) <= '9' {
			s += string(w[p])
			size++
		}
	}
	d := 1
	num := 0
	for i := size - 1; i >= 0; i-- {
		num += d * int(s[i]-'0')
		d *= 10
	}
	if ind == '0' || ind == '+' {
		return num
	} else {
		return -num
	}

}

// func main() {
// 	s := "12345"
// 	s2 := "str123ing45"
// 	s3 := "012 345"
// 	s4 := "Hello World!"
// 	s5 := "sd+x1fa2W3s4"
// 	s6 := "sd-x1fa2W3s4"
// 	s7 := "sdx1-fa2W3s4"

// 	n := TrimAtoi(s)
// 	n2 := TrimAtoi(s2)
// 	n3 := TrimAtoi(s3)
// 	n4 := TrimAtoi(s4)
// 	n5 := TrimAtoi(s5)
// 	n6 := TrimAtoi(s6)
// 	n7 := TrimAtoi(s7)

// 	fmt.Println(n)
// 	fmt.Println(n2)
// 	fmt.Println(n3)
// 	fmt.Println(n4)
// 	fmt.Println(n5)
// 	fmt.Println(n6)
// 	fmt.Println(n7)
// }
