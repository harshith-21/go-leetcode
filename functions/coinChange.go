package functions

import "sort"

func CoinChange(coins []int, amount int) int {
	sort.Ints(coins)
	if amount == 0 {
		return 0
	}
	if coins[0] > amount {
		return -1
	}
	for i := len(coins) - 1; i >= 0; i-- {
		amountnow := amount
		count := 0
		for j := i; j >= 0; {
			if amountnow >= coins[j] {
				amountnow -= coins[j]
				count++
			} else {
				j--
			}
			if amountnow == 0 {
				break
			}
		}
		if amountnow == 0 {
			return count
		}
	}
	return -1
}

//func CoinChange(coins []int, amount int) int {
//	count := 0
//	sort.Ints(coins)
//	fmt.Println(coins)
//	if amount == 0 {
//		return 0
//	}
//	if coins[0] > amount {
//		return -1
//	}
//	for i := len(coins) - 1; i >= 0; {
//		fmt.Println("i: ", i)
//		if amount >= coins[i] {
//			amount -= coins[i]
//			count++
//		} else {
//			i--
//		}
//		if amount == 0 {
//			break
//		}
//	}
//	if amount != 0 {
//		return -1
//	}
//	return count
//}
