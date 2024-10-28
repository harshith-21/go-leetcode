package functions

func LeftRightDifference(nums []int) []int {
	// sum := SumOfArray(nums)
	Rsum := func(arr []int) int {
		sum := 0
		for _, x := range arr {
			sum += x
		}
		return sum
	}(nums)
	Lsum := 0
	var ans []int
	for i := 0; i < len(nums); i++ {
		val := Rsum - nums[i] - Lsum
		ans = append(ans, func() int {
			if val > 0 {
				return val
			} else {
				return -1 * val
			}
		}())
		Lsum += nums[i]
		Rsum -= nums[i]
	}
	return ans
}
