package functions

func longestPalindrome(s string) string {
	largestString := ""
	for i := 0; i < len(s)-1; i++ {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				if checkPalindrome(s, i, j) {
					if len(largestString) < (j - i + 1) {
						largestString = s[i:j]
					}
				}
			}
		}
	}
	if len(s) == 1 || len(s) == 0 {
		return s
	}
	if largestString == "" {
		return s[:1]
	}
	return largestString
}

func checkPalindrome(s string, a int, b int) bool {
	for b > a {
		if s[b] != s[a] {
			return false
		}
		b--
		a++
	}
	return true
}
