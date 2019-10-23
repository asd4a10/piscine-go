package IsPrintable

// import (
// 	"fmt"
// )

func IsPrintable(a string) bool {
	for i := range a {
		if a[i] <= 31 {
			return false
		}
	}
	return true
}
