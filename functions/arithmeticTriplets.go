package functions

//2367. Number of Arithmetic Triplets

// V3
func ArithmeticTriplets(nums []int, diff int) int {
	count := 0
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if nums[j]-nums[i] > diff {
				break
			}
			if nums[j]-nums[i] < diff {
				continue
			}
			for k := 0; k < length; k++ {
				if nums[k]-nums[j] > diff {
					break
				}
				if nums[k]-nums[j] < diff {
					continue
				}
				count++
			}
		}
	}
	return count
}

// V2
//func ArithmeticTriplets(nums []int, diff int) int {
//	count := 0
//	length := len(nums)
//	for i := 0; i < length; i++ {
//		for j := 0; j < length; j++ {
//			if nums[j]-nums[i] > diff {
//				break
//			}
//			if nums[j]-nums[i] < diff {
//				continue
//			}
//			if nums[j]-nums[i] == diff {
//				for k := 0; k < length; k++ {
//					if nums[k]-nums[j] > diff {
//						break
//					}
//					if nums[k]-nums[j] < diff {
//						continue
//					}
//					if nums[k]-nums[j] == diff {
//						count++
//					}
//				}
//			}
//		}
//	}
//	return count
//}

// V1
//func ArithmeticTriplets(nums []int, diff int) int {
//	count := 0
//	for i := 0; i < len(nums); i++ {
//		for j := i + 1; j < len(nums); j++ {
//			if nums[j]-nums[i] > diff {
//				break
//			}
//			if nums[j]-nums[i] < diff {
//				continue
//			}
//			for k := j + 1; k < len(nums); k++ {
//				if nums[k]-nums[j] > diff {
//					break
//				}
//				if nums[k]-nums[j] < diff {
//					continue
//				}
//				if nums[k]-nums[j] == diff && nums[j]-nums[i] == diff {
//					count++
//				}
//			}
//		}
//	}
//	return count
//}
