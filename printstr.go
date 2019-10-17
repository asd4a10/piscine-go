package PrintStr

func PrintStr(str string) {

	for _, symbol := range str {
		PrintRune(symbol)
	}
	PrintRune(10)
}
