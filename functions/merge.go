package functions

import "sort"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	sort.Ints(append(nums1[:m-1], nums2...))
}
