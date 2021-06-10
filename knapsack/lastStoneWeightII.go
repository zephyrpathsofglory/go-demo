// 1049. Last Stone Weight II

package knapsack

func LastStoneWeightII(stones []int) int {
	sum := 0
	for _, k := range stones {
		sum += k
	}

	target := sum / 2

	a := make([][]int, len(stones)+1)
	for idx := range a {
		a[idx] = make([]int, target+1)
	}

	for i := 0; i <= len(stones); i++ {
		for j := 0; j <= target; j++ {
			if i == 0 || j == 0 {
				a[i][j] = 0
				continue
			}

			if j < stones[i-1] {
				a[i][j] = a[i-1][j]
			} else {
				a[i][j] = max(a[i-1][j], stones[i-1]+a[i-1][j-stones[i-1]])
			}
		}
	}

	return sum - a[len(stones)][target]*2
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
