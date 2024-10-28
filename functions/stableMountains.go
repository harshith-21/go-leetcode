package functions

func stableMountains(height []int, threshold int) []int {
	count := []int{}
	for i := 1; i < len(height); i++ {
		if height[i] > threshold {
			count = append(count, height[i])
		}
	}
	return count
}
