package FindNextPrime

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	var ans bool = true
	for i := 2; i <= nb/2; i++ {
		if nb%i == 0 {
			ans = false
			break
		}
	}
	return ans
}

func FindNextPrime(nb int) int {
	for i := nb; ; i++ {
		if IsPrime(i) == true {
			return i
		}
	}
}
