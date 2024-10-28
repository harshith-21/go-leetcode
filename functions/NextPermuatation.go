package functions

import (
	"fmt"
	"math"
	"sort"
)

func NextPermutation(nums []int) {
	target := -1
	for i := len(nums) - 1; i >= 1; i-- {
		if nums[i-1] < nums[i] {
			target = i - 1
			break
		}
	}
	fmt.Println(target)
	if target == -1 {
		sort.Ints(nums)
	} else {
		nextmaxnumindex := target
		mindiff := math.MaxInt8
		for i := target + 1; i < len(nums); i++ {
			if nums[i] > nums[target] && nums[i]-nums[nextmaxnumindex] < mindiff {
				nextmaxnumindex = i
				mindiff = nums[nextmaxnumindex] - nums[target]
			}
		}
		temp := nums[nextmaxnumindex]
		nums[nextmaxnumindex] = nums[target]
		nums[target] = temp

		sort.Ints(nums[target+1:])
	}

}
