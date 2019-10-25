package main

import (
	// "github.com/01-edu/z01"
	"os"
	"fmt"
)

func Index(s string, toFind string) int {
	if toFind == "" {
		return 0
	}
	size := 0
	for range toFind {
		size++
	}
	b := -1
	count := 0
	for i := range s {
		if count == size {
			break
		} else if s[i] == toFind[count] {
			if count == 0 {
				b = i
			}
			count++
		} else {
			count = 0
			b = -1
		}
	}
	if count != size {
		return -1
	}
	if b == -1 {
		return -1
	}
	return b
}
func Concat(str1 string, str2 string) string {
	w := str1 + str2
	return w
}

func main() {
	name := os.Args
	s1 = 0
	for range name {
		s1++
	}
	if Index(name[1], "--help") != -1 || Index(name[1], "-h") != -1 {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("    This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("This flag will behave like a boolean, if it is called it will order the argument.")
		return
	} else if s1 == 1 {
		fmt.Println("--insert")
		fmt.Println("  -i")
		fmt.Println("    This flag inserts the string into the string passed as argument.")
		fmt.Println("--order")
		fmt.Println("  -o")
		fmt.Println("This flag will behave like a boolean, if it is called it will order the argument.")
	}
	ord := 0
	u := 0
	if Index(name[1], "--insert") != -1 {
		u = 1
		s := name[1]
		var w string
		for i := range s {
			if i > 8 {
				w += string(s[i])
			}
		}
		fmt.Println(w)
	} else if Index(name[1], "-i") != -1 {
		u = 1
		s := name[1]
		var w string
		for i := range s {
			if i > 2 {
				w += string(s[i])
			}
		}
	}
	if Index(name[1], "--order") != -1 || Index(name[1], "-o") != -1 {
		ord = 1
	} else if Index(name[2], "--order") != -1 {
		ord = 1
	}
	if u == 1 {
		if flag[:9] == "--insert="
	}
	
	
	
	
	
	
	// count := 1
	// for count == 1 {
	// 	for i := range name {
	// 		if i == size-1 {
	// 			break
	// 		} else if name[i] > name[i+1] {
	// 			a = name[i+1]
	// 			name[i+1] = name[i]
	// 			name[i] = a
	// 		}
	// 	}
	// 	count = 0
	// 	for i := range name {
	// 		if i == size-1 {
	// 			break
	// 		} else if name[i] > name[i+1] {
	// 			count = 1
	// 		}
	// 	}
	// }
}