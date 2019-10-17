package main

import (
    "fmt"
)

func UltimateDivMod(a *int, b *int) {
	c := *a
	d := *b
	*a = c / d
	*b = c % d
}


func main() {
	a := 15
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
