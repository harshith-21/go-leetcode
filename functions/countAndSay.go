package functions

func CountAndSay(n int) string {
	ans := "1"
	var newans string
	for i := 0; i < n; i++ {
		newans = ""
		for j := 0; j < len(ans)-1; j++ {
			count := 0
			for k := j + 1; k < len(ans); k++ {
				if ans[j] == ans[k] {
					count++
					continue
				} else {
					newans += string(rune(count)) + string(ans[j])
					count = 0
				}
			}
		}
		ans = newans
	}
	return ans
}
