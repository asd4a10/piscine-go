package IterativePower

// import (
//         "fmt"
// )

//1
// package IterativeFactorial

// func IterativeFactorial(nb int) int {
// 	if nb > 20 || nb < 0 {
// 		return 0
// 	} else if nb == 1 || nb == 0 {
// 		return 1
// 	} else {
// 		return IterativeFactorial(nb-1) * nb
// 	}
// }

func IterativePower(nb int, power int) int {
	if power == 0 {
		return nb
	} else if power < 0 {
		return 0
	} else {
	result := 1
	for i := 1; i <= power; i++ {
		result = result * nb
	}
	return result
	}
}
