package main

import "github.com/01-edu/z01"
import "os"


func main() {
	// arguments := os.Args
	name := os.Args[0]
	for i := range name {
		z01.PrintRune(rune(name[i]))
	}
	z01.PrintRune(10)

}
