package functions

import "sort"

func ThreeSum(nums []int) [][]int {
	var ans [][]int

	if len(nums) < 3 {
		return ans
	}
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, len(nums)-1

		for left < right {

			sum := nums[i] + nums[left] + nums[right]

			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[left], nums[right]})

				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}

	}
	return ans
}
