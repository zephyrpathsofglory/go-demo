// 494. Target Sum

package knapsack

func FindTargetSumWays(nums []int, target int) int {
	sum := 0
	for _, k := range nums {
		sum += k
	}

	a := make([][]int, len(nums)+1)
	for idx := range a {
		a[idx] = make([]int, sum+1)
	}

	for i := 0; i <= len(nums); i++ {
		for j := 0; j <= sum; j++ {
			if i == 0 {
				if j == 0 {
					a[i][j] = 1
				} else {
					a[i][j] = 0
				}
				continue
			}

			var v1, v2 int
			if j-nums[i-1] < 0 {
				v1 = a[i-1][nums[i-1]-j]
			} else {
				v1 = a[i-1][j-nums[i-1]]
			}

			if j+nums[i-1] > sum {
				v2 = 0
			} else {
				v2 = a[i-1][j+nums[i-1]]
			}
			v := v1 + v2

			if i == len(nums) && j == target {
				return v
			} else {
				a[i][j] = v
			}
		}
	}

	return 0
}
