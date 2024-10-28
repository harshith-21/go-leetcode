package functions

import "fmt"

func frequencySort(nums []int) []int {
	mpp := map[int]int{}

	for i := 0; i < len(nums); i++ {
		mpp[nums[i]]++
	}
	fmt.Println(mpp)

	for j := 0; j < len(nums); j++ {
		
	}

	return nums
}

func TestfrequencySort() {
	nums := []int{1, 2, 3, 4}
	frequencySort(nums)
}
