package functions

import "fmt"

func MajorityElement(nums []int) []int {
	counts := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		_, ok := counts[nums[i]]
		if !ok {
			counts[nums[i]] = 1
			continue
		}
		counts[nums[i]] = counts[nums[i]] + 1
	}

	for key, value := range counts {
		fmt.Println(key, value)
	}

	ans := []int{}
	for v, c := range counts {
		if c > len(nums)/3 {
			ans = append(ans, v)
		}
	}
	return ans
}
