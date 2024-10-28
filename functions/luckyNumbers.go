package functions

//
//func luckyNumbers(matrix [][]int) []int {
//	ans := []int{}
//	for i := 0; i < len(matrix); i++ {
//		minindex := 0
//		for j := 1; j < len(matrix[0]); j++ {
//			if matrix[i][minindex] > matrix[i][j] {
//				minindex = j
//			}
//		}
//		// to check if that min index in row is max index in column
//		maxindex := 0
//		for k := 0; k < len(matrix); k++ {
//			if matrix[maxindex][minindex] < matrix[k][minindex] {
//				maxindex = k
//			}
//		}
//	}
//}
//
////func luckyNumbers(matrix [][]int) []int {
////	mininrows := []int{}
////	maxincol := []int{}
////	for _, arr := range matrix {
////		mini := math.MaxInt64
////		for _, e := range arr {
////			mini = min(mini, e)
////		}
////		mininrows = append(mininrows, mini)
////	}
////
////	for i := 0; i < len(matrix); i++ {
////		maxi := math.MinInt8
////		for j := 0; j < len(matrix[0]); j++ {
////			maxi = max(maxi, matrix[j][i])
////		}
////		maxincol = append(maxincol, maxi)
////	}
////
////	ans := []int{}
////	for _, e1 := range mininrows {
////		for _, e2 := range maxincol {
////			if e1 == e2 {
////				ans = append(ans, e1)
////			}
////		}
////	}
////	fmt.Println(mininrows)
////	fmt.Println(maxincol)
////	return ans
////}
