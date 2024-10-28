package functions

import "math"

func maxProduct(nums []int) int {
	suffix := 1
	prefix := 1
	maxi := math.MinInt64
	leng := len(nums)
	for i := 0; i < leng; i++ {

		if suffix == 0 {
			suffix = 1
		}
		if prefix == 0 {
			prefix = 1
		}

		suffix *= nums[i]
		prefix *= nums[leng-i-1]

		maxi = max(maxi, max(suffix, prefix))

	}
	return maxi
}
