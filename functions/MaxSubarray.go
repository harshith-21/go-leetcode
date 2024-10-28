package functions

func MaxSubArray(nums []int) int {
	var sum int
	var maxSum int
	sum = nums[0]
	maxSum = sum
	for i := 1; i < len(nums); i++ {
		if sum < nums[i] {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		maxSum = max(sum, maxSum)
	}
	return maxSum
}
