package StrLen

func StrLen(str string) int {
	size := 0
	for range str {
		size++
	}
	return size
}
