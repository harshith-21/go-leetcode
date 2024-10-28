package functions

func MaximumWealth(accounts [][]int) int {
	var max int
	for _, x := range accounts {
		if max < SumOfArray(x) {
			max = SumOfArray(x)
		}
	}
	return max
}

func SumOfArray(arr []int) int {
	var sum int
	for _, num := range arr {
		sum += num
	}
	return sum
}
