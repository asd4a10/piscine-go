package PrintNbrInOrder

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var w string
	size := 0
	count := 1
	var a byte
	for n > 0 {
		w += string(rune(n%10 + '0'))
		n /= 10
		size++
	}
	sen := []byte(w)
	for count == 1 {
		for i := range sen {
			if i == size-1 {
				break
			} else if sen[i] > sen[i+1] {
				a = sen[i+1]
				sen[i+1] = sen[i]
				sen[i] = a
			}
		}
		count = 0
		for i := range sen {
			if i == size-1 {
				break
			} else if sen[i] > sen[i+1] {
				count = 1
			}
		}
	}
	b := string(sen)
	for i := range b {
		z01.PrintRune(rune(b[i]))
	}
}

// func main() {
// 	PrintNbrInOrder(32223423434)
// 	PrintNbrInOrder(0)
// 	PrintNbrInOrder(321)
// }
