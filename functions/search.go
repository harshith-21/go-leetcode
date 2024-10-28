package functions

func search(nums []int, target int) bool {
	return BinarySearch(nums, 0, len(nums)-1, target)
}

func findPivot(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			return i
		}
	}
	return -1
}

func BinarySearch(nums []int, a int, b int, target int) bool {
	if nums[a] == target {
		return true
	}
	if nums[b] == target {
		return true
	}
	if nums[(a+b)/2] > target {
		return BinarySearch(nums, a, (a+b)/2, target)
	} else {
		return BinarySearch(nums, (a+b)/2, b, target)
	}
}
