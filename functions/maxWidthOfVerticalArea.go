package functions

import "sort"

func MaxWidthOfVerticalArea(points [][]int) int {
	var firstElementArr []int
	for _, val := range points {
		firstElementArr = append(firstElementArr, val[0])
	}
	sort.Ints(firstElementArr)
	var maxdiff int
	for i := 0; i < (len(firstElementArr) - 1); i++ {
		diff := firstElementArr[i+1] - firstElementArr[i]
		if maxdiff < diff {
			maxdiff = diff
		}
	}
	return maxdiff
}
