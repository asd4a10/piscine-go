package Sqrt

func Sqrt(nb int) int {
	i := 1
	for ; i*i <= nb; i++ {
	}
	if (i-1)*(i-1) == nb {
		return i - 1
	} else {
		return 0
	}
}
