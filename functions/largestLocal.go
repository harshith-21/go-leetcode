package functions

//func LargestLocal(grid [][]int) [][]int {
//	ans := make([][]int, len(grid))
//	for i := 0; i <= (len(grid) - 3); i++ {
//		for j := 0; j <= (len(grid) - 3); j++ {
//			ans[i][j] = maxxx(grid, i, j)
//		}
//	}
//	return ans
//}
//
//func maxxx(grid [][]int, a int, b int) int {
//	return max(grid[a][b], grid[a][b+1], grid[a][b+2],
//		grid[a+1][b], grid[a+1][b+1], grid[a+1][b+2],
//		grid[a+2][b], grid[a+2][b+1], grid[a+2][b+2])
//}

//func LargestLocal(grid [][]int) [][]int {
//	//ans := make([][]int, len(grid)-2)
//	size := len(grid) - 2
//	ans := make([][]int, size)
//	//for i := range ans {
//	//	ans[i] = make([]int, len(grid[i]))
//	//}
//	for i := 0; i <= (len(grid) - 3); i++ {
//		ans[i] = make([]int, len(grid[i]))
//	}
//
//	for i := 0; i <= (len(grid) - 3); i++ {
//		for j := 0; j <= (len(grid[i]) - 3); j++ {
//			ans[i][j] = maxxx(grid, i, j)
//		}
//	}
//	return ans
//}

func largestLocal(grid [][]int) [][]int {
	length := len(grid)
	ans := make([][]int, length-2)
	for i := 0; i < (length - 2); i++ {
		ans[i] = make([]int, length-2)
	}
	for i := 0; i <= (length - 3); i++ {
		for j := 0; j <= (length - 3); j++ {
			ans[i][j] = maxxx(grid, i, j)
		}
	}
	return ans
}

func maxxx(grid [][]int, a int, b int) int {
	return max(grid[a][b], grid[a][b+1], grid[a][b+2],
		grid[a+1][b], grid[a+1][b+1], grid[a+1][b+2],
		grid[a+2][b], grid[a+2][b+1], grid[a+2][b+2])
}
