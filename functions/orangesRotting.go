package functions

// idea : find all fresh apples and do dfs/bfs on it for 4 directions
// if one of the node is 2 then it will be rotten in future (helps in finding isolated apples)
// for all fresh apples (value : 1) height of tree (from fresh node in 4 directions and leaves ending in "2")
// maintain a "visited 0/1" matrix to remove any double iterations

//func orangesRotting(grid [][]int) int {
//	if isAllRotten(grid) {
//		return 0
//	}
//
//}
//
//func isAllRotten(grid [][]int) bool {
//	for i := 0; i < len(grid); i++ {
//		for j := 0; j < len(grid[i]); j++ {
//			if grid[i][j] == 1 {
//				return false
//			}
//		}
//	}
//	return true
//}
//
//func HasIsolatedApple(grid [][]int) bool {
//
//}
//
//func HasSurroundingApple(grid [][]int, i int, j int) bool {
//
//}
