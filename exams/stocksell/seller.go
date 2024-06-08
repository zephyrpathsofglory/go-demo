/*
   1 <= prices.length <= 105
   0 <= prices[i] <= 104
	 选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票(即只能买卖一次)。设计一个算法来计算你所能获取的最大利润。
*/

package stocksell

func MaxProfit(prices []int) (maxProfit int) {
	if len(prices) == 0 {
		panic("prices is empty")
	}

	min := prices[0]

	for i := 1; i < len(prices); i++ {
		iPrice := prices[i]
		if iPrice < min {
			min = iPrice
		} else {
			if iPrice-min > maxProfit {
				maxProfit = iPrice - min
			}
		}
	}

	return
}
