package dp

import "math"

// brute force
// n(n-1)/2 aka O(n^2)
func buySellStock(prices []int) int {
	profit := 0
	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[j]-prices[i] > profit {
				profit = prices[j] - prices[i]
			}
		}
	}
	return profit
}

// BuySellStockConst ...
func BuySellStockConst(prices []int) int {
	// issue converting this to an int
	min := int(math.Inf(-1))
	maxProfit := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i]-min > maxProfit {
			maxProfit = prices[i] - min
		}
	}

	return maxProfit
}
