package IsPrime

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	var ans bool = true
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			ans = false
		}
	}
	return ans
}
