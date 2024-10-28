package functions

func MinimumCost(m int, n int, horizontalCut []int, verticalCut []int) int {
	cost := 0
	hsum := sumOfArray(horizontalCut)
	vsum := sumOfArray(verticalCut)
	if m*hsum > n*vsum {
		cost += vsum
		cost += (m - 1) * hsum
	} else if m == n {
		if vsum > hsum {
			cost += hsum
			cost += vsum * m
		} else {
			cost += vsum
			cost += hsum * m
		}
	} else {
		cost += hsum
		cost += (n - 1) * vsum
	}
	return cost
}

func sumOfArray(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}
