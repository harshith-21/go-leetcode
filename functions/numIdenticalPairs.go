package functions

func NumIdenticalPairs(nums []int) int {
	count := 0
	for i, n := range nums {
		for j := i + 1; j < len(nums); j++ {
			if n == nums[j] {
				count++
			}
		}
	}
	return count
}
