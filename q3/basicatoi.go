package main

import (
	"fmt"
)

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

func main() {
	s := "12345"
	s2 := "0000000012345"
	s3 := "012 345"
	s4 := "Hello World!"

	n := piscine.BasicAtoi2(s)
	n2 := piscine.BasicAtoi2(s2)
	n3 := piscine.BasicAtoi2(s3)
	n4 := piscine.BasicAtoi2(s4)

	fmt.Println(n)
	fmt.Println(n2)
	fmt.Println(n3)
	fmt.Println(n4)

}
