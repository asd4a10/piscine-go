package Compare

// import (
// 	"fmt"
// )

func Compare(a, b string) int {
	sizea := 0
	sizeb := 0
	for range a {
		sizea++
	}
	for range b {
		sizeb++
	}
	c := []byte(a)
	d := []byte(b)
	if sizea > sizeb {
		for i := 0; i < sizeb; i++ {
			if c[i] > d[i] {
				return 1
			} else if c[i] < d[i] {
				return -1
			}
		}
		return 1
	} else if sizea < sizeb {
		for i := 0; i < sizeb; i++ {
			if c[i] > d[i] {
				return 1
			} else if c[i] < d[i] {
				return -1
			}
		}
		return -1
	} else {
		for i := 0; i < sizeb; i++ {
			if c[i] > d[i] {
				return 1
			} else if c[i] < d[i] {
				return -1
			}
		}
		return 0
	}
}

// func main() {
// 	fmt.Println(Compare("Hello!", "Hello!"))
// 	fmt.Println(Compare("Salut!", "lut!"))
// 	fmt.Println(Compare("Ola!", "Ol"))
// }
