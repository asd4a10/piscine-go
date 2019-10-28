package main

import "github.com/01-edu/z01"
import "os"

func printStr(str string) {
	arrayStr := []rune(str)

	for i := 0; i < len(arrayStr); i++ {
		z01.PrintRune(arrayStr[i])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) int {
	if nbr % 2 == 0 {
		return 1
	} else {
		return 0
	}
}

func main() {
	arguments := os.Args
	len := -1 
	for range arguments {
		len++
	}
	if isEven(len) == 1 {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
