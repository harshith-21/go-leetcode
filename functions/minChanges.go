package functions

//
//import (
//	"fmt"
//	"math"
//)
//
//func minChanges(nums []int, k int) int {
//	diff := []int{}
//	for i := 0; i < len(nums)/2; i++ {
//		diff = append(diff, int(math.Abs(float64(nums[i])-float64(nums[len(nums)-i-1]))))
//	}
//	count := map[int]int{}
//	for _, num := range diff {
//		_, ok := count[num]
//		if ok {
//			count[num] += 1
//		} else {
//			count[num] = 1
//		}
//	}
//	maxf := 0
//	var commondiff int
//	for dif, f := range count {
//		if maxf < f {
//			maxf = f
//			commondiff = dif
//		}
//	}
//	for i := 0; i < len(diff); i++ {
//		if diff[i] != commondiff {
//			if diff[i]
//		}
//	}
//
//}
//
//func TestminChanges() {
//	fmt.Println(minChanges([]int{1, 0, 1, 2, 4, 3}, 4))
//	fmt.Println(minChanges([]int{0, 1, 2, 3, 3, 6, 5, 4}, 6))
//}
