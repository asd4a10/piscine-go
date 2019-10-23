package IsLower

// import (
// 	"fmt"
// )

func IsLower(a string) bool {
	for i := range a {
		if a[i] > 'z' || a[i] < 'a' {
			return false
		}
	}
	return true
}
