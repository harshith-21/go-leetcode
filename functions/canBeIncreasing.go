package functions

import "fmt"

func CanBeIncreasing(nums []int) bool {
	highs := []int{}
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			if nums[i] < nums[j] {
				highs = append(highs, j)
				fmt.Println(i, j)
				break
			}
		}
	}
	fmt.Println(highs)
	if len(highs) == 1 {
		return true
	}
	return false
}
