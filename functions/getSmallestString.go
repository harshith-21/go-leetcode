package functions

func GetSmallestString(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes)-1; i++ {
		if same_parity(int(runes[i]), int(runes[i+1])) {
			if int(runes[i]) > int(runes[i+1]) {
				temp := runes[i]
				runes[i] = runes[i+1]
				runes[i+1] = temp
				break
			}
		}
	}
	return string(runes)
}

func same_parity(m int, n int) bool {
	return n%2 == m%2
}
