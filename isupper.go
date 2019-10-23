package IsLower

// import (
// 	"fmt"
// )

func IsLower(a string) bool {
	for i := range a {
		if a[i] > 'Z' || a[i] < 'A' {
			return false
		}
	}
	return true
}
