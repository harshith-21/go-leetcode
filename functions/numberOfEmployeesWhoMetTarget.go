package functions

func NumberOfEmployeesWhoMetTarget(hours []int, target int) int {
	count := 0
	for _, x := range hours {
		if target <= x {
			count++
		}
	}
	return count
}
