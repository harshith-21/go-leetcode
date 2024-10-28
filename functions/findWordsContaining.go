package functions

import "strings"

func FindWordsContaining(words []string, x byte) []int {
	var ans []int
	for i, str := range words {
		if strings.Contains(str, string(x)) {
			ans = append(ans, i)
		}
	}
	return ans
}
