package Index

// import (
// 	"fmt"
// )

func Index(s string, toFind string) int {
	size := 0
	for range toFind {
		size++
	}
	b := -1
	count := 0
	for i := range s {
		if count == size {
			break
		} else if s[i] == toFind[count] {
			if count == 0 {
				b = i
			}
			count++
		} else {
			count = 0
			b = 0
		}
	}
	if b == -1 {
		return -1
	}
	return b

}

// func main() {
// 	fmt.Println(Index("Hello!", "el"))
// 	fmt.Println(Index("Salut!", "alu"))
// 	fmt.Println(Index("Ola!", "hOl"))
// }
