package student

import "github.com/01-edu/z01"

func Raid1a(x, y int) {
	if x > 0 && y > 0 {
		for i := 1; i <= y; i++ {
			for j := 1; j <= x; j++ {
				if i == 1 || i == y {
					if j == 1 || j == x {
						z01.PrintRune('o') //e
					} else {
						z01.PrintRune('-')
					}
				} else {
					if j == 1 || j == x {
						z01.PrintRune('|') //e
					} else {
						z01.PrintRune(' ')
					}
				}
			}
			z01.PrintRune(10)
		}
	} else {

	}
}
