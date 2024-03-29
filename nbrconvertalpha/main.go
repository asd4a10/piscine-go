package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	u := 0
	for index, i := range name {
		if index == 0 {
			continue
		}
		if index == 1 && i == "--upper" {
			u = 1
			z01.PrintRune(' ')
			continue
		}
		a := 0
		for _, j := range i {
			if rune(j) < '0' || rune(j) > '9' {
				break
			}
			a = a*10 + int(j) - 48
		}
		if a > 26 || a < 1 {
			z01.PrintRune(' ')
		} else if u == 1 {
			z01.PrintRune(rune(a + 64))
		} else {
			z01.PrintRune(rune(a + 96))
		}
	}
	z01.PrintRune(10)
}
