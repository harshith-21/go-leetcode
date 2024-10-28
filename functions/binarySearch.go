package functions

func binarySearch(nums []int, target int) int {
	return binaryse(nums, target, 0, len(nums)-1)
}

func binaryse(nums []int, target int, a int, b int) int {
	mid := a + (b-a)/2

	if nums[a] != target && nums[b] != target && b-a == 1 {
		return -1
	}

	if target == nums[mid] {
		return mid
	} else if target > nums[mid] {
		a = mid + 1
		return binaryse(nums, target, a, b)
	} else {
		b = mid - 1
		return binaryse(nums, target, a, b)
	}
}
