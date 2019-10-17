package PrintStr

func PrintStr(str string) {

	for _, symbol := range str {
		z01.PrintRune(symbol)
	}
	z01.PrintRune(10)
}
