package functions

import "fmt"

func BuildArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		ans[i] = nums[nums[i]]
	}
	return ans
}

func BuildArray2(nums []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		fmt.Print(nums[nums[i]], " ")
		ans = append(ans, nums[nums[i]])
	}
	return ans
}
