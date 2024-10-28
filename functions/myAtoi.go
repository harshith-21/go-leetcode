package functions

import "math"

func MyAtoi(s string) int {
	i, sum, negative := 0, 0, false
	for s[i] == ' ' {
		i++
	}
	if s[i] == '-' {
		negative = true
	} else if s[i] == '+' {
		negative = false
	} else {
		negative = false
	}
	i++
	for i < len(s) {
		if int(s[i])-48 >= 0 && int(s[i])-48 <= 9 {

			if sum > (math.MaxInt32-int(s[i])+48)/10 {
				if negative {
					return math.MinInt32
				}
				return math.MaxInt32
			}

			sum = sum*10 + int(s[i]) - 48
		} else {
			return sum
		}
		i++
	}
	if negative {
		sum *= -1
	}
	return sum
}
