package functions

import "fmt"

func minimumLength(s string) int {
	count := map[rune]int{}
	for _, letter := range s {
		_, ok := count[letter]
		if ok {
			count[letter] += 1
		} else {
			count[letter] = 1
		}
	}
	fmt.Println(count)
	minlength := 0
	for _, counts := range count {
		// 1 : 1
		// 2 : 2
		// 3 : 1
		// 4 : 2
		// 5 : 1
		// 6 : 2
		// 7 : 1
		if counts%2 == 0 {
			minlength += 2
		} else {
			minlength += 1
		}

	}
	return minlength
}

func TestminimumLength() {
	fmt.Println(minimumLength("abaacbcbb"))
	fmt.Println(minimumLength("aa"))
}
