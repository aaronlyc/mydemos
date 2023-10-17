package algo

import "math"

func bestTimeBuySellStock(prices []int) int {
	if len(prices) <= 0 || len(prices) > 100000 {
		return 0
	}

	minPrice := math.MaxInt
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < 0 || prices[i] > 10000 {
			return 0
		}
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		}
	}

	return maxProfit
}
