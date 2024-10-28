package functions

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxcount := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
		} else {
			count = 0
		}
		maxcount = max(maxcount, count)
	}
	return maxcount
}
