package main

import (
	"github.com/01-edu/z01"
	"os"
)

func main() {
	name := os.Args
	size := 0
	for range name {
		size++
	}
	count := 1
	a := ""
	for count == 1 {
		for i := range name {
			if i == size-1 {
				break
			} else if name[i] > name[i+1] {
				a = name[i+1]
				name[i+1] = name[i]
				name[i] = a
			}
		}
		count = 0
		for i := range name {
			if i == size-1 {
				break
			} else if name[i] > name[i+1] {
				count = 1
			}
		}
	}
	for j := 0; j < size - 1; j++ {
		a := []rune(name[j])
		for i := range name[j] {
			z01.PrintRune(a[i])
		}
		z01.PrintRune(10)
	}
}
