package od

/*
机器人走迷宫 深度优先搜索

1.房间由XY的方格组成，例如下图为24的大小。每一个方格以坐标(x，y)描述

2.机器人固定从方格(0，0)出发，只能向东或者向北前进。出口固定为房间的最东北角，如下图的方格(5，3)。用例保证机器人可以从入口走到出口。

3.房间有些方格是墙壁，如(4，1)，机器人不能经过那儿。

4.有些地方是一旦到达就无法走到出口的，如标记为B的方格，称之为陷阱方格。

5.有些地方是机器人无法到达的的，如标记为A的方格，称之为不可达方格，不可达方格不包括墙壁所在的位置.

6.如下示例图中，陷阱方格有2个，不可达方格有3个。

7.请为该机器人实现路径规划功能: 给定房间大小、墙壁位置，请计算出陷阱方格与不可达方格分别有多少个


输入描述

第一行为房间的X和Y (0 < X，Y = 1000)
第二行为房间中墙壁的个数N (0 = N< X*Y)
同一行中如果有多个数据以一个空格隔开，用例保证所有的输入数据均合法。 (结尾不带回车换行)

输出描述

陷阱方格与不可达方格数量，两个信息在一行中输出，以一个空格隔开。(结尾不带回车换行)

-------------------------
有图片说明，查看 assets 文件夹
*/

func mazeWalker(x, y int, wall [][2]int) (trapCount, unreachableCount int) {
	matrix := make([][]int, 0, x)
	for i := 0; i < x; i++ {
		matrix = append(matrix, make([]int, y))
	}

	/*
			0  默认值 不可到达
			1  墙
		  -1 陷进
			2  可到达
	*/

	for _, w := range wall {
		matrix[w[0]][w[1]] = 1
	}
	matrix[x-1][y-1] = 2

	var f func(int, int) bool
	f = func(i, j int) bool {
		if i >= x || j >= y {
			return false
		}

		if matrix[i][j] == 1 {
			return false
		}

		if matrix[i][j] == -1 {
			return false
		}

		if matrix[i][j] == 2 {
			return true
		}

		if matrix[i][j] == 0 {
			// 这里千万不能用 f(i+1, j) || f(i, j+1) 因为这是个短路运算，前面如果为 true，后面不会执行了
			// if f(i+1, j) || f(i, j+1) {
			east := f(i+1, j)
			north := f(i, j+1)
			if east || north {
				matrix[i][j] = 2
			} else {
				matrix[i][j] = -1
			}
		}

		return matrix[i][j] == 2
	}

	f(0, 0)

	for _, i := range matrix {
		for _, j := range i {
			if j == 0 {
				unreachableCount += 1
			} else if j == -1 {
				trapCount += 1
			}
		}
	}

	return
}
