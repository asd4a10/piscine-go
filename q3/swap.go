package StrRev

func StrRev(s string) string {
	var w string
	size := 0
	for range s{
		size++
	}
	for i, _ := range s{
		w += string(s[size - 1 - i])
	}
	return w
}

