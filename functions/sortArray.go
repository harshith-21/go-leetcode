package functions

import "fmt"

func sortArray(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	firstArray := sortArray(nums[:len(nums)/2])
	secondArray := sortArray(nums[len(nums)/2:])
	return merge(firstArray, secondArray)
}

func merge(nums1 []int, nums2 []int) []int {
	finalarray := []int{}
	i := 0
	j := 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] > nums2[j] {
			finalarray = append(finalarray, nums2[j])
			j++
		} else {
			finalarray = append(finalarray, nums1[i])
			i++
		}
	}
	for ; i < len(nums1); i++ {
		finalarray = append(finalarray, nums1[i])
	}
	for ; j < len(nums2); j++ {
		finalarray = append(finalarray, nums2[j])
	}
	fmt.Println(finalarray)
	return finalarray
}

func TestMergesort() {
	nums := []int{4, 7, 3, 8, 0, 1, 3, 5, 2, 9}
	fmt.Println(sortArray(nums))
}
