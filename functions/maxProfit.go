package functions

func maxProfit(prices []int) int {
	minprice, maxProfitval := 0, 0
	for _, price := range prices {
		if maxProfitval < price-minprice {
			maxProfitval = price - minprice
		}
		if minprice > price {
			minprice = price
		}
	}
	return maxProfitval
}
