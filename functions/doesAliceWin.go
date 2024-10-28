package functions

import "fmt"

func doesAliceWin(s string) bool {
	count := 0
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' {
		}
		count++
	}
	fmt.Println(count)
	if count > 0 {
		return true
	}
	return false
}
