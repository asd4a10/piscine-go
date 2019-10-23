package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[0]
	a := []rune(name)
	for i := range a {
		z01.PrintRune(a[i])
	}
	z01.PrintRune(10)

}
