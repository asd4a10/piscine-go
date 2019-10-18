package main

import "github.com/01-edu/z01"

func PrintStr(str string) {
	
	for _, symbol := range str{
		z01.PrintRune(symbol)
	}
	z01.PrintRune(10)
}


func main() {
	str := "Hello World!"
	PrintStr(str)
}
