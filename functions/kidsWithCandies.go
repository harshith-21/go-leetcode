package functions

import "slices"

func KidsWithCandies(candies []int, extraCandies int) []bool {
	maxcandies := slices.Max(candies)
	var ans []bool
	for _, x := range candies {
		if (maxcandies - x) <= extraCandies {
			ans = append(ans, true)
		} else {
			ans = append(ans, false)
		}
	}
	return ans
}
