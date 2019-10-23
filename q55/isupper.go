package IsUpper

// import (
// 	"fmt"
// )

func IsUpper(a string) bool {
	for i := range a {
		if a[i] > 'Z' || a[i] < 'A' {
			return false
		}
	}
	return true
}
