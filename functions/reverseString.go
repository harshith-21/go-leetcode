package functions

func reverseString(s []byte) {
	var buff byte
	for i := 0; i < len(s)/2; i++ {
		buff = s[i]
		s[i] = s[len(s)-i-1]
		s[len(s)-i-1] = buff
	}
}
