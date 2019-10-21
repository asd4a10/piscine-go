package IterativeFactorial

// import (
//         "fmt"
// )

//1
// package IterativeFactorial

func IterativeFactorial(nb int) int {
	if nb > 20 || nb < 0 {
		return 0
	} else if nb == 1 || nb == 0 {
		return 1
	} else {
		return IterativeFactorial(nb-1) * nb
	}
}

//1
// func IterativePower(nb int, power int) int {
// 	result := 1
// 	for i := 1; i <= power ; i++ {
// 		result = result * nb
// 	}
// 	return result
// }

// 3
// func IterativePower(nb int, power int) int {
// 	if power == 1 {
// 		return nb
// 	} else {
// 		return nb * IterativePower(nb, power - 1)
// 	}
// }

// 5
// func Fibonacci(index int) int {
// 	if index == 0 {
// 		return 0
// 	} else if index == 1 {
// 		return 1
// 	} else if index < 0 {
// 		return -1
// 	} else {
// 		return Fibonacci(index - 1) + Fibonacci(index - 2)
// 	}
// }

// else {
// 	st := nb
// 	for ;Sqrt(nb - 1) == 0; nb-- {
// 	}
// 	if (Sqrt(nb - 1) + 1) * (Sqrt(nb - 1) + 1) > st {
// 		return 0
// 	} else {
// 		return Sqrt(nb - 1) + 1
// 	}
// }

// 6 sqrt
// func Sqrt(nb int) int {
// 	i := 1
// 	for ; i * i <= nb; i++ {
// 	}
// 	if (i - 1) * (i - 1) == nb {
// 		return i - 1;
// 	} else {
// 		return 0
// 	}
// }

//7 prime
// func IsPrime(nb int) bool {
// 	var ans bool = true
// 	for i := 2; i < nb; i++ {
// 		if nb % i == 0 {
// 			ans = false
// 		}
// 	}
// 	return ans
// }

//8 Findnext
// func IsPrime(nb int) bool {
// 	var ans bool = true
// 	for i := 2; i < nb; i++ {
// 		if nb % i == 0 {
// 			ans = false
// 		}
// 	}
// 	return ans
// }

// func FindNextPrime(nb int) int {
// 	for i := nb; ; i++ {
// 		if IsPrime(i) == true {
// 			return i
// 		}
// 	}
// }

// 9

// func main() {

// }
