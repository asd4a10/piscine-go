package RecursiveFactorial

// import (
//         "fmt"
// )

//1
// package IterativeFactorial

func RecursiveFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		return RecursiveFactorial(nb-1) * nb
	}
}
