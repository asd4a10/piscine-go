package BasicJoin

// import (
// 	"fmt"
// )

func BasicJoin(s []string) string {
	var w string
	for i := range s {
		w += s[i]
	}
	return w
}

// func main() {
// 	toConcat := []string{"Hello!", " How", " are", " you?"}
// 	fmt.Println(BasicJoin(toConcat))
// }
