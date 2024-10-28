package functions

// 2828. Check if a String Is an Acronym of Words

func IsAcronym(words []string, s string) bool {
	if len(s) != len(words) {
		return false
	}
	for i, word := range words {
		if word[0] != s[i] {
			return false
		}
	}
	return true
}
