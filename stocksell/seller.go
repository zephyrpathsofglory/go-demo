/*
   1 <= prices.length <= 105
   0 <= prices[i] <= 104
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
