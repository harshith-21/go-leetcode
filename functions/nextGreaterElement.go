package functions

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	ans := []int{}
	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for k := j; k < len(nums2); k++ {
					if nums2[k] > nums1[i] {
						ans = append(ans, nums2[k])
						break
					}
					if k == len(nums2)-1 {
						ans = append(ans, -1)
					}
				}
			}
		}
	}
	return ans
}
