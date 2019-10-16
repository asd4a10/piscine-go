package main

import "github.com/01-edu/z01"

func main() {

	for i := 97; i < 123; {
		z01.PrintRune(rune(i))
		i++
	}
	z01.PrintRune(run(10))
}
