package SortIntegerTable

// import (
// 	"fmt"
// )

func SortIntegerTable(table []int) {
	size := len(table)
	count := 1
	a := 0
	for ; count == 1 ; {
		for i := range table {
			if i == size - 1 {
				break
			} else if table[i] > table[i + 1] {
				a = table[i + 1]
				table[i + 1] = table[i]
				table[i] = a
			}
		}  
			count = 0
		for i := range table {
			if i == size - 1 {
				break
			} else if table[i] > table[i + 1] {
				count = 1
			}
		}
	}
}


// func main() {
// 	s := []int{0,4,3,7,1,0}
// 	SortIntegerTable(s)
// 	fmt.Println(s)
// }
