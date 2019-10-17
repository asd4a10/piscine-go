package StrLen

func StrLen(str string) int {
	size := 1
	for  index := range str {
		size = index
	}
	size++
	return size
}
