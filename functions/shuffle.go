package functions

func Shuffle(nums []int, n int) []int {
	ans := []int{}
	for i := 0; i < n; i++ {
		ans = append(ans, nums[i], nums[i+n])
	}
	return ans
}
