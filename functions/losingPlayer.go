package functions

import "fmt"

func losingPlayer(x int, y int) string {
	turns := min(x, y/4)
	if turns%2 == 0 {
		return "Bob"
	}
	return "Alice"
}

func TestlosingPlayer() {
	fmt.Println(losingPlayer(2, 7))
	fmt.Println(losingPlayer(4, 11))
}
