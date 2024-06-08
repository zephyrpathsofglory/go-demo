// 1049. Last Stone Weight II

/*
有一堆石头，用整数数组 stones 表示。其中 stones[i] 表示第 i 块石头的重量。

每一回合，从中选出任意两块石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：

如果 x == y，那么两块石头都会被完全粉碎；
如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
最后，最多只会剩下一块 石头。返回此石头 最小的可能重量 。如果没有石头剩下，就返回 0。



示例 1：

输入：stones = [2,7,4,1,8,1]
输出：1
解释：
组合 2 和 4，得到 2，所以数组转化为 [2,7,1,8,1]，
组合 7 和 8，得到 1，所以数组转化为 [2,1,1,1]，
组合 2 和 1，得到 1，所以数组转化为 [1,1,1]，
组合 1 和 1，得到 0，所以数组转化为 [1]，这就是最优值。
示例 2：

输入：stones = [31,26,33,21,40]
输出：5


提示：

1 <= stones.length <= 30
1 <= stones[i] <= 100
*/

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
