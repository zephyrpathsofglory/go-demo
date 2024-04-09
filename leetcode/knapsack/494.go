// 494. Target Sum

/*
给你一个非负整数数组 nums 和一个整数 target 。

向数组中的每个整数前添加 '+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：

例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。



示例 1：

输入：nums = [1,1,1,1,1], target = 3
输出：5
解释：一共有 5 种方法让最终目标和为 3 。
-1 + 1 + 1 + 1 + 1 = 3
+1 - 1 + 1 + 1 + 1 = 3
+1 + 1 - 1 + 1 + 1 = 3
+1 + 1 + 1 - 1 + 1 = 3
+1 + 1 + 1 + 1 - 1 = 3
示例 2：

输入：nums = [1], target = 1
输出：1


提示：

1 <= nums.length <= 20
0 <= nums[i] <= 1000
0 <= sum(nums[i]) <= 1000
-1000 <= target <= 1000
*/

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
