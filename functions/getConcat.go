package functions

func GetConcatenation(nums []int) []int {
	return append(nums, nums...)
}
