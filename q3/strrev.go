
package main

import (
	"fmt"
)

func StrRev(s string) string {
	var w string
	size := 0
	for range s {
		size++
	}
	for i := range s {
		w = w + string(s[i])
	}
	return w
}

func main() {
	s := "Hello World!"
	s = StrRev(s)
	fmt.Println(s)
}

