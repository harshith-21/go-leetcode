package functions

func CreateTargetArray(nums []int, index []int) []int {
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		ans = InsertAtPositionInArray(index[i], nums[i], ans)
	}
	return ans
}

func InsertAtPositionInArray(pos int, val int, arr []int) []int {
	return append(arr[:pos], append([]int{val}, arr[pos:]...)...)
}
