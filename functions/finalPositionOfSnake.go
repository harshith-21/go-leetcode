package functions

func finalPositionOfSnake(n int, commands []string) int {
	x , y := 0, 0
	for _, command := range commands {
		if command == "UP" {
			y++
		} else if command == "DOWN"  {
			y--
		} else if command == "LEFT" {
			x--
		} else {
			x++
		}
	}
	return x *n + y
}