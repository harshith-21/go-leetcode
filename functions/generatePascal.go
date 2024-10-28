package functions

func Generate(numRows int) [][]int {
	ans := [][]int{{1}}
	for i := 1; i < numRows; i++ {
		ans = append(ans, generateNext(ans[i-1]))
	}
	return ans
}

func generateNext(nums []int) []int {
	newnums := append([]int{0}, nums...)
	newnums = append(newnums, 0)
	ans := []int{}
	for i := 0; i < len(newnums)+1; i++ {
		ans = append(ans, newnums[i]+newnums[i+1])
	}
	return ans
}
