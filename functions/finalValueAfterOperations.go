package functions

func FinalValueAfterOperations(operations []string) int {
	count := 0
	for _, str := range operations {
		if str == "X++" || str == "++X" {
			count++
		}
		if str == "X--" || str == "--X" {
			count--
		}
	}
	return count
}
