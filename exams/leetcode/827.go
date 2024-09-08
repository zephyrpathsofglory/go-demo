/*
给你一个大小为 n x n 二进制矩阵 grid 。最多 只能将一格 0 变成 1 。

返回执行此操作后，grid 中最大的岛屿面积是多少？

岛屿 由一组上、下、左、右四个方向相连的 1 形成。



示例 1:

输入: grid = [[1, 0], [0, 1]]
输出: 3
解释: 将一格0变成1，最终连通两个小岛得到面积为 3 的岛屿。
示例 2:

输入: grid = [[1, 1], [1, 0]]
输出: 4
解释: 将一格0变成1，岛屿的面积扩大为 4。
示例 3:

输入: grid = [[1, 1], [1, 1]]
输出: 4
解释: 没有0可以让我们变成1，面积依然为 4。


提示：

n == grid.length
n == grid[i].length
1 <= n <= 500
grid[i][j] 为 0 或 1
*/

package leetcode

func largestIsland(grid [][]int) int {
	var dfs func(int, int, int, int) int
	dimen := len(grid)
	max := 0

	numberCountMap := make(map[int]int)

	dfs = func(i, j, number, count int) int {
		if (i < 0 || i >= dimen) || (j < 0 || j >= dimen) || grid[i][j] != 1 {
			return count
		}

		count++
		grid[i][j] = number
		count = dfs(i-1, j, number, count)
		count = dfs(i+1, j, number, count)
		count = dfs(i, j-1, number, count)
		count = dfs(i, j+1, number, count)

		numberCountMap[number] = count
		return count
	}

	number := 2
	for i := range grid {
		for j := range grid {
			if grid[i][j] == 1 {
				count := dfs(i, j, number, 0)
				if count > max {
					max = count
				}
				number++
			}
		}
	}

	checkNeighbor := func(i, j int) (int, int) {
		if i >= 0 && i < dimen && j >= 0 && j < dimen {
			count, ok := numberCountMap[grid[i][j]]
			if ok {
				return grid[i][j], count
			}
		}

		return 0, 0
	}

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				numberUp, count := checkNeighbor(i-1, j)
				size := count

				numberDown, count := checkNeighbor(i+1, j)
				if numberDown != numberUp {
					size += count
				}
				numberLeft, count := checkNeighbor(i, j-1)
				if numberLeft != numberDown && numberLeft != numberUp {
					size += count
				}
				numberRight, count := checkNeighbor(i, j+1)
				if numberRight != numberDown && numberRight != numberLeft && numberRight != numberUp {
					size += count
				}

				size += 1

				if size > max {
					max = size
				}
			}
		}
	}

	return max
}
