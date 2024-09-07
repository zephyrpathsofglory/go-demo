/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ["1","1","1","1","0"],
  ["1","1","0","1","0"],
  ["1","1","0","0","0"],
  ["0","0","0","0","0"]
]
输出：1
示例 2：

输入：grid = [
  ["1","1","0","0","0"],
  ["1","1","0","0","0"],
  ["0","0","1","0","0"],
  ["0","0","0","1","1"]
]
输出：3


提示：

m == grid.length
n == grid[i].length
1 <= m, n <= 300
grid[i][j] 的值为 '0' 或 '1'

*/

package leetcode

func numIslands(grid [][]byte) int {
	rows := len(grid)
	cols := len(grid[0])
	var dfs func(int, int)

	dfs = func(i1, i2 int) {
		if !(0 <= i1 && i1 < rows) || !(0 <= i2 && i2 < cols) || grid[i1][i2] != byte('1') {
			return
		}

		grid[i1][i2] = byte('2')

		dfs(i1-1, i2)
		dfs(i1+1, i2)
		dfs(i1, i2-1)
		dfs(i1, i2+1)

		return
	}

	count := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == byte('1') {
				dfs(i, j)
				count++
			}
		}
	}

	return count
}
