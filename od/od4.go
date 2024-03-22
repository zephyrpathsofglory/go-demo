package od

import (
	"cmp"
	"slices"
)

/*
题目描述
假设书本的叠放有这样的规则，当A书的长度和宽度都大于B书时，可以将其B书置于A的上方，堆叠摆放，请设计一个程序，根据输入的书本长宽，
计算最多可以堆叠摆放多少本书？

输入
[[16,15], [13, 12], [15, 14]]

输出
3

说明
这里代表有3本书，第1本长宽分别为16和15，第2本长宽为13和12，第3本长宽为15和14，它们可以按照 [13, 12],[15, 14],[16,15] 的顺序堆叠，
所以输出3
*/

// 动态规划，实质是求最长递增子序列，但本题有两个因素需要考虑，可以通过排序将长度置为有序，这样其实就
// 是对宽度求最长递增子序列，且可能存在长度相同的情况，在更新dp数组判定时，也要考虑到
func bookStack(books [][2]int) (maxHeight int) {
	dp := make([]int, len(books), len(books))
	for i := range dp {
		dp[i] = 1
	}

	maxHeight = 1

	slices.SortFunc(books, func(i, j [2]int) int {
		if i[0] < j[0] {
			return cmp.Compare(i[0], j[0])
		} else {
			return cmp.Compare(i[1], j[1])
		}
	})

	for i := 1; i < len(books); i++ {
		bookI := books[i]
		for j := 0; j < i; j++ {
			if bookI[0] > books[j][0] && bookI[1] > books[j][1] {
				dp[i] = Max(dp[i], dp[j]+1)
			}
		}

		maxHeight = Max(dp[i], maxHeight)
	}

	return
}

// Max returns the larger of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}
