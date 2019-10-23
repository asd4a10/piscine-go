package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args[0]
	for i := range name {
		z01.PrintRune(rune(name[i]))
	}
	z01.PrintRune(10)

}
